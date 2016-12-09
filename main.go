package main

import (
	"net/http"

	"github.com/egreb/websitestatus/broker"
	"github.com/egreb/websitestatus/handlers"
	"github.com/egreb/websitestatus/worker"
	"github.com/julienschmidt/httprouter"
)

// set routes

func main() {
	var router = httprouter.New()
	broker := broker.NewServer()
	router.NotFound = http.FileServer(http.Dir("app/static"))
	router.GET("/", handlers.IndexHandler)
	router.GET("/status/all", handlers.StatusHandler)
	router.GET("/events", broker.ServeHTTP)
	worker.Worker(broker)
	http.ListenAndServe(":8080", router)
}
