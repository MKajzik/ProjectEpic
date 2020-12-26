package slack

// //CreateGameBlocks export
// func CreateGameBlocks(text string, image string, url string, emoji bool) WebhookJSON {
// 	var msg WebhookJSON
// 	var value string = "click_me_123"
// 	var action string = "button-action"

// 	msg.Blocks = make([]Blocks, 3)

// 	msg.Blocks[0].Type = "section"
// 	msg.Blocks[0].Text.Type = "mrkdwn"
// 	msg.Blocks[0].Text.Text = "Siema, dzisiaj Epic zaserwował nam nową darmową grę. Poniżej sprawdźcie ją i nie zapomnijcie jej *ODEBRAĆ*!"
// 	msg.Blocks[1].Type = "section"
// 	msg.Blocks[1].Text.Type = "mrkdwn"
// 	msg.Blocks[1].Text.Text = text
// 	msg.Blocks[1].Accessory = &Accessory{}
// 	msg.Blocks[1].Accessory.Type = "image"
// 	msg.Blocks[1].Accessory.ImageURL = &image
// 	msg.Blocks[1].Accessory.AltText = &text
// 	msg.Blocks[2].Type = "section"
// 	msg.Blocks[2].Text.Type = "mrkdwn"
// 	msg.Blocks[2].Text.Text = "Odbierz mnie pliska. *NO PLISKA*"
// 	msg.Blocks[2].Accessory = &Accessory{}
// 	msg.Blocks[2].Accessory.Type = "button"
// 	msg.Blocks[2].Accessory.Text = &Text{}
// 	msg.Blocks[2].Accessory.Text.Type = "plain_text"
// 	msg.Blocks[2].Accessory.Text.Text = "ODBIERZ"
// 	msg.Blocks[2].Accessory.Text.Emoji = &emoji
// 	msg.Blocks[2].Accessory.Value = &value
// 	msg.Blocks[2].Accessory.URL = &url
// 	msg.Blocks[2].Accessory.ActionID = &action

// 	return msg

// }

// //CreateNoGameBlock export
// func CreateNoGameBlock() WebhookJSON {
// 	var msg WebhookJSON
// 	msg.Blocks = make([]Blocks, 1)

// 	msg.Blocks[0].Type = "section"
// 	msg.Blocks[0].Text.Type = "mrkdwn"
// 	msg.Blocks[0].Text.Text = "Dzisiaj nie ma zadnej gry do odebrania. Sorki :P!"

// 	return msg
// }
