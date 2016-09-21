package site

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/egreb/sitesstatus/utils"
	"github.com/julienschmidt/httprouter"
)

type Status struct {
	Site         Site      `json: site`
	Status       int       `json: status`
	ResponseTime float64   `json: responsetime`
	Updated      time.Time `json: updated`
}

func getStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	site := GetSite(ps.ByName("name"))

	status := status(site)
	j, err := json.Marshal(status)
	utils.Check(err)
	w.Write(j)
}

func status(site Site) Status {
	t := time.Now()
	res, err := http.Get(site.Url)
	utils.Check(err)
	s := Status{
		Site: Site{
			Id:   site.Id,
			Name: site.Name,
			Url:  site.Url,
		},
		Status:       res.StatusCode,
		ResponseTime: utils.TimeDurationInSeconds(t),
		Updated:      t,
	}
	return s
}

func GetStatuses() []Status {
	sites := ReadFile()
	statuses := make([]Status, len(sites))
	for i := 0; i < len(sites); i++ {
		st := status(sites[i])
		statuses[i] = st
	}
	return statuses
}
