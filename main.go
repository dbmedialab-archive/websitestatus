package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"net/http"
	"time"
)

var router = httprouter.New()

// set routes
func init() {
	router.GET("/sites", getSites)
	router.GET("/status/:name", getStatus)
}

func main() {
	http.ListenAndServe(":8080", router)
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

func getSites(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sites := readFile()
	j, err := json.Marshal(sites)
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
			Name: site.Name,
			Url:  site.Url,
		},
		Status:       res.StatusCode,
		ResponseTime: timeDurationInSeconds(t),
	}
	return s
}

func timeDurationInSeconds(t time.Time) float64 {
	return time.Since(t).Seconds()

}

type Site struct {
	Name string `json: name`
	Url  string `json: url`
}

type Status struct {
	Site         Site    `json: site`
	Status       int     `json: status`
	ResponseTime float64 `json: responsetime`
}