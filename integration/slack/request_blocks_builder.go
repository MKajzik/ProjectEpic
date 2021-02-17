package slack

//Block export
type block struct {
	Type      string     `json:"type"`
	Text      text       `json:"text"`
	Accessory *accessory `json:"accessory,omitempty"`
}

//BlockBuilder export
type BlockBuilder interface {
	SetType(string) BlockBuilder
	SetText(text) BlockBuilder
	SetAccessory(accessory) BlockBuilder
	Build() block
	Reset()
}

type blockBuilder struct {
	Type            string
	Text            text
	AccessoryOption *accessory
}

//SetType export
func (b *blockBuilder) SetType(s string) BlockBuilder {
	b.Type = s
	return b
}

//SetText export
func (b *blockBuilder) SetText(t text) BlockBuilder {
	b.Text = t
	return b
}

//SetAccessory export
func (b *blockBuilder) SetAccessory(a accessory) BlockBuilder {
	b.AccessoryOption = &a
	return b
}

//Build export
func (b *blockBuilder) Build() block {
	return block{
		Type:      b.Type,
		Text:      b.Text,
		Accessory: b.AccessoryOption,
	}
}

//Reset export
func (b *blockBuilder) Reset() {
	b.Type = ""
	b.Text = text{}
	b.AccessoryOption = nil

}

//NewBlockBuilder export
func NewBlockBuilder() BlockBuilder {
	return &blockBuilder{}
}
