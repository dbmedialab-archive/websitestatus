package handlers

import (
	"net/http"
	"text/template"

	"github.com/egreb/sitesstatus/site"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	data := site.Site{}
	t, _ := template.ParseFiles("./static/index.html")
	t.Execute(w, data)
}
