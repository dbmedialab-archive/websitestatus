package worker

import (
	"encoding/json"
	"log"
	"time"

	"github.com/egreb/websitestatus/broker"
	"github.com/egreb/websitestatus/site"
	"github.com/egreb/websitestatus/utils"
)

func Worker(broker *broker.Broker) {
	go func() {
		for {
			time.Sleep(time.Second * 5)
			status := site.GetStatuses()
			j, err := json.Marshal(status)
			utils.Check(err)
			log.Println("Receiving event")
			broker.Notifier <- j
		}
	}()
}
