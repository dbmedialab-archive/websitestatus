package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/egreb/sitesstatus/site"
)

func SiteHandler(w http.ResponseWriter, r *http.Request) {
	sitesJSON, _ := json.Marshal(site.ReadFile())
	w.Header().Set("Content-Type", "application/json")
	w.Write(sitesJSON)
}
