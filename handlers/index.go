package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/egreb/sitesstatus/utils"
	"github.com/egreb/websitestatus/site"
	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	data := site.Site{}
	t, err := template.ParseFiles("./static/index.html")
	utils.Check(err)
	t.Execute(w, data)
}

func StatusHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	statuses := site.GetStatuses()
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(statuses)
	if err != nil {
		fmt.Fprint(w, "Could not get statuses")
	}
	w.Write(j)
}

func SiteHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sitesJSON, _ := json.Marshal(site.ReadSitesFromFile())
	w.Header().Set("Content-Type", "application/json")
	w.Write(sitesJSON)
}
