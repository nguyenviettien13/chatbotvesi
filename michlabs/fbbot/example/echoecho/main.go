package main

import (
	"github.com/nguyenviettien13/chatbotvesi/michlabs/fbbot"
)

const PORT int = 8080
const VERIFYTOKEN string = "your_secure_token"
const PAGEACCESSTOKEN string = "your_beloved_page_access_token"

func main() {
	var e EchoEcho

	bot := fbbot.New(PORT, VERIFYTOKEN, PAGEACCESSTOKEN)
	bot.AddMessageHandler(e)
	bot.Run()
}

type EchoEcho struct{}

func (e EchoEcho) HandleMessage(bot *fbbot.Bot, msg *fbbot.Message) {
	// Echo... echo...
	m := fbbot.NewTextMessage(msg.Text)
	bot.Send(msg.Sender, m)
}
