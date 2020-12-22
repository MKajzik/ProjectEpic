package slack

//WebhookJSON export
type WebhookJSON struct {
	Blocks []Blocks `json:"blocks"`
}

//Text export
type Text struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji *bool  `json:"emoji,omitempty"`
}

//Accessory export
type Accessory struct {
	Type     string  `json:"type"`
	Text     *Text   `json:"text,omitempty"`
	ImageURL *string  `json:"image_url,omitempty"`
	AltText  *string `json:"alt_text,omitempty"`
	Value    *string `json:"value,omitempty"`
	URL      *string `json:"url,omitempty"`
	ActionID *string `json:"action_id,omitempty"`
}

//Blocks export
type Blocks struct {
	Type      string     `json:"type"`
	Text      Text       `json:"text"`
	Accessory *Accessory `json:"accessory,omitempty"`
}
