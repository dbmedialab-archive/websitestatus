package status

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"sync"

	"github.com/egreb/websitestatus/site"
)

var state = make(map[string]CurrentStatus)
var mutex = &sync.Mutex{}

// Check the status for every site, if status is bad send notification to slack channel
func Check(s []site.Status) {
	for _, x := range s {
		fmt.Println(x.Site.Name, ":", state[x.Site.Url].Response.Count)
		mutex.Lock()
		oldStatus := state[x.Site.Url]
		currentStatus := oldStatus

		currentStatus.Status = true
		if x.Status != 200 {
			currentStatus.Status = false
		}

		// If responsetime is higher then 5 seconds
		if int(x.ResponseTime) >= 5000 {
			currentStatus.Response.Count++
		} else {
			currentStatus.Response.Count = 0
		}

		if currentStatus.Response.Count >= 10 {
			currentStatus.Response.Err = true
		} else {
			currentStatus.Response.Err = false
		}
		state[x.Site.Url] = currentStatus
		mutex.Unlock()

		attachments := Attachments{}
		sendSlack := false
		if (oldStatus.Status != currentStatus.Status || oldStatus.Response.Err != currentStatus.Response.Err) && currentStatus.Status && !currentStatus.Response.Err {
			attachments.Text = fmt.Sprintf("<%s|%s> kjører igjen som normalt\n", x.Site.Url, x.Site.Name)
			attachments.Color = "#36a64f"
			sendSlack = true
		} else {
			if currentStatus.Response.Err && currentStatus.Response.Count%10 == 0 {
				fmt.Println("Check")

				if currentStatus.Response.Count > 10 {
					attachments.Text = fmt.Sprintf("<!channel> <%s|%s> opplever fortsatt veldig lang lastetid, vennligst kontroller at nettisden fungerer som den skal", x.Site.Url, x.Site.Name)
				} else {
					attachments.Text = fmt.Sprintf("<!channel> <%s|%s> opplever veldig lang lastetid, vennligst kontroller at nettisden fungerer som den skal", x.Site.Url, x.Site.Name)
				}
				attachments.Color = "#c00"
				sendSlack = true
			}
		}

		if sendSlack {
			sr := SlackResponse{[]Attachments{attachments}}
			SendMessage(sr)
		}
	}
}

// SendMessage to slack channel
func SendMessage(response SlackResponse) {
	slackURL := os.Getenv("SLACK_URL")
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

// SlackResponse are being sent to the slack channel
type SlackResponse struct {
	Attachments []Attachments `json:"attachments"`
}

// CurrentStatus structure
type CurrentStatus struct {
	Status   bool         `json:"status"`
	Response ResponseTime `json:"responsetime"`
}

// ResponseTime structure
type ResponseTime struct {
	Err   bool `json:"error"`
	Count int  `json:"count"`
}

// Attachments structure
type Attachments struct {
	Text  string `json:"text"`
	Color string `json:"color"`
}
