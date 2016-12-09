package worker

import (
	"encoding/json"
	"log"
	"time"

	"github.com/dbmedialab/websitestatus/broker"
	"github.com/dbmedialab/websitestatus/controllers"
	"github.com/dbmedialab/websitestatus/utils"
)

// Worker checks websites every 5 seconds and returns the json feed
func Worker(broker *broker.Broker) {
	go func() {
		for {
			time.Sleep(time.Second * 5)
			s := controllers.GetStatuses()
			controllers.ControlStatus(s)
			j, err := json.Marshal(s)
			utils.Ok(err)
			log.Println("Receiving event")
			broker.Notifier <- j
		}
	}()
}
