package controllers

import (
	"io/ioutil"
	"net/http"
	"time"

	"github.com/egreb/websitestatus/models"
	"github.com/egreb/websitestatus/utils"
)

// GetStatus returns the status of the site
func GetStatus(site models.Site) models.Status {
	t := time.Now()
	res, err := http.Get(site.URL)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	s := models.Status{
		Site: models.Site{
			ID:   site.ID,
			Name: site.Name,
			URL:  site.URL,
		},
		Status:       0,
		Size:         0,
		ResponseTime: 0,
		Updated:      utils.DateToString(t),
		Error:        "",
	}
	if err != nil {
		s.Error = "Could not connect to website"
	} else {
		s.Status = res.StatusCode
		s.Size = float64(len(body) / 1000)
		s.ResponseTime = utils.TimeDurationInMilliseconds(t) // cast to milliseconds
	}
	return s
}

// GetStatuses returns statuses for all sites in the json file
func GetStatuses() []models.Status {
	sites := ReadSitesFromFile()
	statusArr := make([]models.Status, len(sites))

	for i, x := range sites {
		statusArr[i] = GetStatus(x)
	}

	return statusArr
}
