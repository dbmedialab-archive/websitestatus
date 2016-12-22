package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/dbmedialab/websitestatus/models"
)

var state = make(map[string]models.CurrentStatus)
var mutex = &sync.Mutex{}

const statusOK = 200
const maxResponseTime = 3000
const countTreshold = 10

// ControlStatus checks the status for every site, if status is bad send notification to slack channel
func ControlStatus(sites []models.Status) {
	for _, site := range sites {
		mutex.Lock()
		oldStatus := state[site.Site.URL]
		currentStatus := updateStatus(oldStatus, site)
		state[site.Site.URL] = currentStatus
		mutex.Unlock()

		buildSlackMessage(oldStatus, currentStatus, site)
	}
}

func buildSlackMessage(old models.CurrentStatus, curr models.CurrentStatus, site models.Status) {
	attachments := models.Attachments{}
	send := false

	if old.Error != curr.Error {
		if !curr.Error {
			attachments.Text = fmt.Sprintf("<%s|%s> kjører igjen som normalt", site.Site.URL, site.Site.Name)
			attachments.Color = "#36a64f"
			send = true
		} else {
			attachments.Text = fmt.Sprintf("<!channel> <%s|%s> ser ut til å være nede, får ikke respons 200. Vennligst kontroller", site.Site.URL, site.Site.Name)
			attachments.Color = "#c00"
		}
	}

	if !send {
		if old.Response.Error != curr.Response.Error {
			if !curr.Response.Error {
				attachments.Text = fmt.Sprintf("<%s|%s> kjører igjen som normalt etter å ha opplevd en periode med lang lastetid", site.Site.URL, site.Site.Name)
				attachments.Color = "#36a64f"
			} else {
				attachments.Text = fmt.Sprintf("<!channel> <%s|%s> opplever veldig lang lastetid, vennligst kontroller at nettsiden fungerer som den skal", site.Site.URL, site.Site.Name)
				attachments.Color = "#c00"
			}
			send = true
		}
	}

	if send {
		// Build slackmessage
		message := models.SlackMessage{[]models.Attachments{attachments}}
		// Send message to slack channel
		Broadcast(message)
	}
}

func updateStatus(current models.CurrentStatus, site models.Status) models.CurrentStatus {
	current.Error = false
	if site.Status != statusOK {
		current.Error = true
	}

	// If responsetime is higher then 5 seconds increase count
	// Else reset counter
	if int(site.ResponseTime) >= maxResponseTime {
		current.Response.Count++
	} else {
		current.Response.Count = 0
	}

	current.Response.Error = false
	if current.Response.Count >= countTreshold {
		current.Response.Error = true
	}
	return current
}

// Broadcast to slack channel
func Broadcast(message models.SlackMessage) {
	slackURL := os.Getenv("SLACK_URL")

	// Check if the SLACK_URL env variable is set, if not messages are not being sent to slack
	if slackURL != "" {
		// parse to json
		j, err := json.Marshal(message)
		if err != nil {
			panic(err.Error)
		}
		req, err := http.NewRequest("POST", slackURL, bytes.NewBuffer(j))
		req.Header.Set("Content-Type", "application/json")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			fmt.Println(err.Error())
		}
		defer resp.Body.Close()
	}
}
