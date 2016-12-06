package models

// SlackResponse are being sent to the slack channel
type SlackResponse struct {
	Attachments []Attachments `json:"attachments"`
}

// CurrentStatus structure
type CurrentStatus struct {
	Status   bool         `json:"status"`
	Response ResponseTime `json:"responsetime"`
}

// ResponseTime structure
type ResponseTime struct {
	Err   bool `json:"error"`
	Count int  `json:"count"`
}

// Attachments structure
type Attachments struct {
	Text  string `json:"text"`
	Color string `json:"color"`
}
