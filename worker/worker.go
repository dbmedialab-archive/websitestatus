package worker

import (
	"encoding/json"
	"log"
	"time"

	"github.com/egreb/websitestatus/broker"
	"github.com/egreb/websitestatus/site"
	"github.com/egreb/websitestatus/status"
	"github.com/egreb/websitestatus/utils"
)

// Worker checks websites every 5 seconds and returns the json feed
func Worker(broker *broker.Broker) {
	go func() {
		for {
			time.Sleep(time.Second * 5)
			s := site.GetStatuses()
			status.Check(s)
			j, err := json.Marshal(s)
			utils.Check(err)
			log.Println("Receiving event")
			broker.Notifier <- j
		}
	}()
}
