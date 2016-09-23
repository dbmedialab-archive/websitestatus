package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"

	"github.com/egreb/websitestatus/site"
	"github.com/julienschmidt/httprouter"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := site.Site{}
	t, _ := template.ParseFiles("./static/index.html")
	t.Execute(w, data)
}

func StatusHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	statuses := site.GetStatuses()
	j, err := json.Marshal(statuses)
	if err != nil {
		fmt.Fprint(w, "Could not get statuses")
	}
	w.Write(j)
}
