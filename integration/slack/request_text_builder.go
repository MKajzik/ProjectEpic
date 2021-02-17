package slack

type text struct {
	Type  string `json:"type"`
	Text  string `json:"text"`
	Emoji *bool  `json:"emoji,omitempty"`
}

//TextBuilder export
type TextBuilder interface {
	SetType(string) TextBuilder
	SetText(string) TextBuilder
	SetEmoji(bool) TextBuilder
	Build() text
	Reset()
}

type textBuilder struct {
	Type        string
	Text        string
	EmojiOption *bool
}

//SetType export
func (t *textBuilder) SetType(s string) TextBuilder {
	t.Type = s
	return t
}

//SetText export
func (t *textBuilder) SetText(s string) TextBuilder {
	t.Text = s
	return t
}

//SetEmoji export
func (t *textBuilder) SetEmoji(b bool) TextBuilder {
	t.EmojiOption = &b
	return t
}

//Build export
func (t *textBuilder) Build() text {
	return text{
		Type:  t.Type,
		Text:  t.Text,
		Emoji: t.EmojiOption,
	}
}

//Reset export
func (t *textBuilder) Reset() {
	t.Type = ""
	t.Text = ""
	t.EmojiOption = nil
}

//NewTextBuilder export
func NewTextBuilder() TextBuilder {
	return &textBuilder{}
}
