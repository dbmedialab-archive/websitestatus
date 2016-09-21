package worker

import (
	"encoding/json"
	"log"
	"time"

	"github.com/egreb/sitesstatus/broker"
	"github.com/egreb/sitesstatus/site"
	"github.com/egreb/sitesstatus/utils"
)

func Worker(broker *broker.Broker) {
	go func() {
		for {
			time.Sleep(time.Second * 5)
			status := site.ReadFile()
			j, err := json.Marshal(status)
			utils.Check(err)
			log.Println("Receiving event")
			broker.Notifier <- j
		}
	}()
}
