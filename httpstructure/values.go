package httpvalues

//Slack export
type Slack struct {
	Blocks []Blocks `json:"blocks"`
}

//Text export
type Text struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

//Accessory export
type Accessory struct {
	Type     string `json:"type"`
	ImageURL string `json:"image_url"`
	AltText  string `json:"alt_text"`
}

//Blocks export
type Blocks struct {
	Type      string    `json:"type"`
	Text      Text      `json:"text"`
	Accessory Accessory `json:"accessory,omitempty"`
}
