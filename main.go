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
	_, _, err := c.PostMessage(CHANNEL, slack.MsgOptionText("Hello World", true))
	if err != nil {
		panic(err)
	}
	fmt.Println("Hello World")
}
