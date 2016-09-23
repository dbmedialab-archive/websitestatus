package site

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/egreb/websitestatus/utils"
	"github.com/julienschmidt/httprouter"
)

type Site struct {
	Id   int    `json: id`
	Name string `json: name`
	Url  string `json: url`
}

func GetSite(name string) Site {
	sites := ReadSitesFromFile()
	var s Site
	for i := 0; i < len(sites); i++ {
		if sites[i].Name == name {
			s := sites[i]
			return s
		}
	}
	return s
}

func ReadSitesFromFile() []Site {
	dat, err := ioutil.ReadFile("sites.json")
	utils.Check(err)
	var sites []Site
	err = json.Unmarshal(dat, &sites)
	utils.Check(err)

	return sites
}

func GetSites(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	sites := ReadSitesFromFile()
	j, err := json.Marshal(sites)
	utils.Check(err)

	w.Write(j)
}
