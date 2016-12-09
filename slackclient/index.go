package slackclient

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/dbmedialab/websitestatus/models"

	"encoding/json"
)

// SendMessage to slack channel
func SendMessage(response models.SlackResponse) {
	const url string = "https://hooks.slack.com/services/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"
	j, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err.Error)
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(j))
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error)
	}
	defer resp.Body.Close()

	fmt.Println("Response status", resp.Status)
	fmt.Println("Headers: ", resp.Header)
	fmt.Println("Body:", resp.Body)
}
