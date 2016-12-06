package models

// Status structure represents the status of a website
type Status struct {
	Site         Site    `json:"site"`
	Status       int     `json:"status"`
	Size         float64 `json:"size"`
	ResponseTime float64 `json:"responsetime"`
	Updated      string  `json:"updated"`
	Error        string  `json:"error"`
}
