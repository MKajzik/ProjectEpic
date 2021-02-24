package slack

//CreateTextBlock exported
func CreateTextBlock(text string) block {
	blockBuilder := NewBlockBuilder()
	textBuilder := NewTextBuilder()

	block := blockBuilder.
		SetType("section").
		SetText(textBuilder.
			SetType("mrkdwn").
			SetText(text).
			Build()).
		Build()

	return block
}

//CreateImageBlock exported
func CreateImageBlock(text string, image string) block {
	blockBuilder := NewBlockBuilder()
	textBuilder := NewTextBuilder()
	accessoryBuilder := NewAccessoryBuilder()

	block := blockBuilder.
		SetType("section").
		SetText(textBuilder.
			SetType("mrkdwn").
			SetText(text).
			Build()).
		SetAccessory(accessoryBuilder.
			SetType("image").
			SetImageURL(image).
			SetAltText(text).
			Build()).
		Build()

	return block
}

//CreateButtonBlock exported
func CreateButtonBlock(text string, buttonText string, url string) block {
	blockBuilder := NewBlockBuilder()
	textBuilder := NewTextBuilder()
	accessoryBuilder := NewAccessoryBuilder()

	block := blockBuilder.
		SetType("section").
		SetText(textBuilder.
			SetType("mrkdwn").
			SetText(text).
			Build()).
		SetAccessory(accessoryBuilder.
			SetType("button").
			SetText(textBuilder.
				SetType("plain_text").
				SetText(buttonText).
				SetEmoji(true).
				Build()).
			SetValue("click_me_123").
			SetURL(url).
			SetActionID("button-action").
			Build()).
		Build()

	return block
}
