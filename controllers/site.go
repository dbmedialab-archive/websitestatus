package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/dbmedialab/websitestatus/models"
	"github.com/dbmedialab/websitestatus/utils"
	"github.com/julienschmidt/httprouter"
)

// GetSite returns site according to string
func GetSite(name string) models.Site {
	sites := ReadSitesFromFile()
	var s models.Site
	for i := 0; i < len(sites); i++ {
		if sites[i].Name == name {
			s := sites[i]
			return s
		}
	}
	return s
}

// ReadSitesFromFile returns all entered sites in the json file
func ReadSitesFromFile() []models.Site {
	dat, err := ioutil.ReadFile("sites.json")
	utils.Ok(err)
	var sites []models.Site
	err = json.Unmarshal(dat, &sites)
	utils.Ok(err)

	return sites
}

// GetSites returns all sites from json file
func GetSites(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sites := ReadSitesFromFile()
	j, err := json.Marshal(sites)
	utils.Ok(err)

	w.Write(j)
}
