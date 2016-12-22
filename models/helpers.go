package models

// SlackMessage are being sent to the slack channel
type SlackMessage struct {
	Attachments []Attachments `json:"attachments"`
}

// CurrentStatus structure
type CurrentStatus struct {
	Error    bool         `json:"status"`
	Response ResponseTime `json:"responsetime"`
}

// ResponseTime structure
type ResponseTime struct {
	Error bool `json:"error"`
	Count int  `json:"count"`
}

// Attachments structure
type Attachments struct {
	Text  string `json:"text"`
	Color string `json:"color"`
}
