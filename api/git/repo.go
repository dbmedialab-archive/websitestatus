package git

import (
	"log"
	"net/http"
)

// GetAll repos
func GetAll() {
	req, err := http.NewRequest("GET", "https://api.github.com/egreb/repos", nil)
	if err != nil {
		log.Fatal(err)
	}
	req.SetBasicAuth("52d0cee116491b528ba9857d84fd0074f155a23f", "x-oauth-basic")

	// res, err := http.DefaultClient.Do(req)
}
