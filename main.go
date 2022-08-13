package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {
	TOKEN := os.Getenv("TOKEN")
	CHANNEL := os.Getenv("CHANNEL")

	c := slack.New(TOKEN)

	params := &slack.GetConversationHistoryParameters{
		ChannelID: CHANNEL,
		// Cursor:             "",
		// Inclusive:          false,
		// Latest:             "",
		// Limit:              5,
		// Oldest:             "",
		// IncludeAllMetadata: false,
	}

	res, err := c.GetConversationHistory(params)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%v", res)
	// _, _, err = c.PostMessage(CHANNEL, slack.MsgOptionText("Hello World", true))
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Println("Hello World")
}
