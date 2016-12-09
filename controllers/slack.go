package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/egreb/websitestatus/models"
)

var state = make(map[string]models.CurrentStatus)
var mutex = &sync.Mutex{}

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

func buildSlackMessage(oldStatus models.CurrentStatus, currentStatus models.CurrentStatus, site models.Status) {
	attachments := models.Attachments{}
	sendSlack := false
	if (oldStatus.Status != currentStatus.Status || oldStatus.Response.Err != currentStatus.Response.Err) && currentStatus.Status && !currentStatus.Response.Err {
		attachments.Text = fmt.Sprintf("<%s|%s> kjører igjen som normalt\n", site.Site.URL, site.Site.Name)
		attachments.Color = "#36a64f"
		sendSlack = true
	} else {
		if currentStatus.Response.Err && currentStatus.Response.Count%10 == 0 {
			fmt.Println("Check")

			if currentStatus.Response.Count > 10 {
				attachments.Text = fmt.Sprintf("<!channel> <%s|%s> opplever fortsatt veldig lang lastetid, vennligst kontroller at nettisden fungerer som den skal", site.Site.URL, site.Site.Name)
			} else {
				attachments.Text = fmt.Sprintf("<!channel> <%s|%s> opplever veldig lang lastetid, vennligst kontroller at nettisden fungerer som den skal", site.Site.URL, site.Site.Name)
			}
			attachments.Color = "#c00"
			sendSlack = true
		}
	}
	if sendSlack {
		sr := models.SlackResponse{[]models.Attachments{attachments}}
		SendMessage(sr)
	}
}

func updateStatus(cs models.CurrentStatus, site models.Status) models.CurrentStatus {
	cs.Status = true
	if site.Status != 200 {
		cs.Status = false
	}

	// If responsetime is higher then 5 seconds
	if int(site.ResponseTime) >= 5000 {
		cs.Response.Count++
	} else {
		cs.Response.Count = 0
	}

	if cs.Response.Count >= 10 {
		cs.Response.Err = true
	} else {
		cs.Response.Err = false
	}

	return cs
}

// SendMessage to slack channel
func SendMessage(response models.SlackResponse) {
	slackURL := os.Getenv("SLACK_URL")

	// Check if the env variable is set, if not exit program
	if slackURL == "" {
		fmt.Println("SLACK_URL er ikke definert!!")
		fmt.Println("Programmet avsluttes")
		os.Exit(0)
	}
	j, err := json.Marshal(response)
	if err != nil {
		panic(err.Error)
	}
	req, err := http.NewRequest("POST", slackURL, bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err.Error)
	}
	defer resp.Body.Close()
}
