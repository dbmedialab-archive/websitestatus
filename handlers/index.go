package handlers

import (
	"encoding/json"
	"net/http"
	"text/template"

	"github.com/egreb/websitestatus/controllers"
	"github.com/egreb/websitestatus/models"
	"github.com/egreb/websitestatus/utils"
	"github.com/julienschmidt/httprouter"
)

// IndexHandler returns all sites
func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := models.Site{}
	t, err := template.ParseFiles("./app/static/index.html")
	utils.Ok(err)
	t.Execute(w, data)
}

// StatusHandler returns all sites statuses
func StatusHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	statuses := controllers.GetStatuses()
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(statuses)
	utils.Ok(err)
	w.Write(j)
}

// SiteHandler returns all sites from json file
func SiteHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sitesJSON, _ := json.Marshal(controllers.ReadSitesFromFile())
	w.Header().Set("Content-Type", "application/json")
	w.Write(sitesJSON)
}
