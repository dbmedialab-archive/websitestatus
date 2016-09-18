package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
	"time"

	"github.com/julienschmidt/httprouter"
)

var router = httprouter.New()

// set routes
/*func init() {
	var broker = NewServer()
	router.NotFound = http.FileServer(http.Dir("static"))
	router.GET("/", index)
	router.GET("/sites", getSites)
	router.Handler("GET", "/statuses", broker)
	router.HandlerFunc("GET", "/status", func(w http.ResponseWriter, r *http.Request) {
		for {
			time.Sleep(time.Second * 2)
			s := statuses()
			j, err := json.Marshal(s)
			check(err)
			broker.Notifier <- []byte(j)
		}
	})
}*/

func main() {
	broker := NewServer()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			s := statuses()
			j, err := json.Marshal(s)
			check(err)
			broker.Notifier <- []byte(j)
		}
	}()

	http.ListenAndServe(":8080", broker)
}

func readFile() []Site {
	dat, err := ioutil.ReadFile("sites.json")
	check(err)
	var sites []Site
	err = json.Unmarshal(dat, &sites)
	check(err)
	return sites
}

func getSite(name string) Site {
	sites := readFile()
	var s Site
	for i := 0; i < len(sites); i++ {
		if sites[i].Name == name {
			s := sites[i]
			return s
		}
	}
	return s
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := Site{}
	t, err := template.ParseFiles("./static/index.html")
	check(err)
	t.Execute(w, data)
}

func getSites(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sites := readFile()
	j, err := json.Marshal(sites)
	check(err)

	w.Write(j)
}

func statuses() []Status {
	sites := readFile()
	statuses := make([]Status, len(sites))
	for i := 0; i < len(sites); i++ {
		st := status(sites[i])
		statuses[i] = st
	}
	return statuses
}

func getStatuses(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := statuses()
	j, err := json.Marshal(s)
	check(err)
	w.Write(j)
}

func check(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func getStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	site := getSite(ps.ByName("name"))

	status := status(site)
	j, err := json.Marshal(status)
	check(err)
	w.Write(j)
}

func status(site Site) Status {
	t := time.Now()
	res, err := http.Get(site.Url)
	check(err)
	s := Status{
		Site: Site{
			Id:   site.Id,
			Name: site.Name,
			Url:  site.Url,
		},
		Status:       res.StatusCode,
		ResponseTime: timeDurationInSeconds(t),
		Updated:      t,
	}
	return s
}

func timeDurationInSeconds(t time.Time) float64 {
	return time.Since(t).Seconds()

}

type Site struct {
	Id   int    `json: id`
	Name string `json: name`
	Url  string `json: url`
}

type Status struct {
	Site         Site      `json: site`
	Status       int       `json: status`
	ResponseTime float64   `json: responsetime`
	Updated      time.Time `json: updated`
}

type Broker struct {
	Notifier       chan []byte
	newClients     chan chan []byte
	closingClients chan chan []byte
	clients        map[chan []byte]bool
}

func NewServer() (broker *Broker) {
	broker = &Broker{
		Notifier:       make(chan []byte, 1),
		newClients:     make(chan chan []byte),
		closingClients: make(chan chan []byte),
		clients:        make(map[chan []byte]bool),
	}

	go broker.listen()

	return
}

func (broker *Broker) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	flusher, ok := rw.(http.Flusher)
	if !ok {
		http.Error(rw, "Streaming unsupported!", http.StatusInternalServerError)
		return
	}

	rw.Header().Set("Content-Type", "text/event-stream")
	rw.Header().Set("Cache-Control", "no-cache")
	rw.Header().Set("Connection", "keep-alive")
	rw.Header().Set("Access-Control-Allow-Origin", "*")

	messageChan := make(chan []byte)

	broker.newClients <- messageChan

	defer func() {
		broker.closingClients <- messageChan
	}()

	notify := rw.(http.CloseNotifier).CloseNotify()

	go func() {
		<-notify
		broker.closingClients <- messageChan
	}()

	for {
		fmt.Fprint(rw, "data: %s\n\n", <-messageChan)

		flusher.Flush()
	}
}

func (broker *Broker) listen() {
	for {
		select {
		case s := <-broker.newClients:

			broker.clients[s] = true
			log.Printf("Client added. %d registered clients", len(broker.clients))

		case s := <-broker.closingClients:
			delete(broker.clients, s)
			log.Printf("Client removed. %d registered clients", len(broker.clients))

		case event := <-broker.Notifier:
			for clientMessageChan, _ := range broker.clients {
				clientMessageChan <- event
			}
		}
	}
}
