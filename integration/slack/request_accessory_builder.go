package slack

//Accessory export
type accessory struct {
	Type     string  `json:"type"`
	Text     *text   `json:"text,omitempty"`
	ImageURL *string `json:"image_url,omitempty"`
	AltText  *string `json:"alt_text,omitempty"`
	Value    *string `json:"value,omitempty"`
	URL      *string `json:"url,omitempty"`
	ActionID *string `json:"action_id,omitempty"`
}

//AccessoryBuilder export
type AccessoryBuilder interface {
	SetType(string) AccessoryBuilder
	SetText(text) AccessoryBuilder
	SetImageURL(string) AccessoryBuilder
	SetAltText(string) AccessoryBuilder
	SetValue(string) AccessoryBuilder
	SetURL(string) AccessoryBuilder
	SetActionID(string) AccessoryBuilder
	Build() accessory
	Reset()
}

type accessoryBuilder struct {
	Type     string
	Text     *text
	ImageURL *string
	AltText  *string
	Value    *string
	URL      *string
	ActionID *string
}

//SetType export
func (a *accessoryBuilder) SetType(s string) AccessoryBuilder {
	a.Type = s
	return a
}

//SetText export
func (a *accessoryBuilder) SetText(t text) AccessoryBuilder {
	a.Text = &t
	return a
}

//SetImageURL export
func (a *accessoryBuilder) SetImageURL(s string) AccessoryBuilder {
	a.ImageURL = &s
	return a
}

//SetAltText export
func (a *accessoryBuilder) SetAltText(s string) AccessoryBuilder {
	a.AltText = &s
	return a
}

//SetValue export
func (a *accessoryBuilder) SetValue(s string) AccessoryBuilder {
	a.Value = &s
	return a
}

//SetURL export
func (a *accessoryBuilder) SetURL(s string) AccessoryBuilder {
	a.URL = &s
	return a
}

//SetActionID export
func (a *accessoryBuilder) SetActionID(s string) AccessoryBuilder {
	a.ActionID = &s
	return a
}

//Build export
func (a *accessoryBuilder) Build() accessory {
	return accessory{
		Type:     a.Type,
		Text:     a.Text,
		ImageURL: a.ImageURL,
		AltText:  a.AltText,
		Value:    a.Value,
		URL:      a.URL,
		ActionID: a.ActionID,
	}
}

//Reset export
func (a *accessoryBuilder) Reset() {
	a.Type = ""
	a.Text = nil
	a.ImageURL = nil
	a.AltText = nil
	a.Value = nil
	a.URL = nil
	a.ActionID = nil
}

//NewAccessoryBuilder export
func NewAccessoryBuilder() AccessoryBuilder {
	return &accessoryBuilder{}
}
