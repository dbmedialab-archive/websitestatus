package site

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/egreb/websitestatus/utils"
)

type Status struct {
	Site         Site    `json: site`
	Status       int     `json: status`
	Size         float64 `json: size`
	ResponseTime float64 `json: responsetime`
	Updated      string  `json: updated`
	Error        string  `json: error`
}

/*func getStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	site := GetSite(ps.ByName("name"))

	status := status(site)
	j, err := json.Marshal(status)
	utils.Check(err)
	w.Write(j)
}*/

func status(site Site) Status {
	t := time.Now()
	res, err := http.Get(site.Url)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	s := Status{
		Site: Site{
			Id:   site.Id,
			Name: site.Name,
			Url:  site.Url,
		},
		Status:       0,
		Size:         0,
		ResponseTime: 0,
		Updated:      DateToString(t),
		Error:        "",
	}
	if err != nil {
		s.Error = "Could not connect to website"
	} else {
		s.Status = res.StatusCode
		s.Size = float64(len(body) / 1000)
		s.ResponseTime = utils.TimeDurationInSeconds(t)
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
