package slack

//WebhookJSON export
type WebhookJSON struct {
	Blocks []Blocks `json:"blocks"`
}

//Text export
type Text struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji *bool  `json:"emoji"`
}

//Accessory export
type Accessory struct {
	Type     string  `json:"type"`
	Text     *Text   `json:"text"`
	ImageURL string  `json:"image_url"`
	AltText  *string `json:"alt_text"`
	Value    *string `json:"value"`
	URL      *string `json:"url"`
	ActionID *string `json:"action_id"`
}

//Blocks export
type Blocks struct {
	Type      string     `json:"type"`
	Text      Text       `json:"text"`
	Accessory *Accessory `json:"accessory,omitempty"`
}
