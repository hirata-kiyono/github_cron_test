package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/slack-go/slack"
)

func main() {
	TOKEN := os.Getenv("TOKEN")
	CHANNEL := os.Getenv("CHANNEL")

	TOKEN = "xoxb-682881597796-3959390218208-vIsjZgpNFLSW9RsvLqD1nHPe"
	CHANNEL = "C029K2CNJ56"

	c := slack.New(TOKEN)
	yesterday := int(time.Now().Add(time.Hour * 24 * (-1)).Unix())
	params := &slack.GetConversationHistoryParameters{
		ChannelID:          CHANNEL,
		Cursor:             "",
		Inclusive:          true,
		Latest:             "",
		Limit:              100,
		Oldest:             strconv.Itoa(yesterday),
		IncludeAllMetadata: false,
	}

	res, err := c.GetConversationHistory(params)
	if err != nil {
		panic(err)
	}
	// _, _, err = c.PostMessage(CHANNEL, slack.MsgOptionText("Hello World", true))
	// if err != nil {
	// 	panic(err)
	// }
	messages := res.Messages
	for i, message := range messages {
		text := fmt.Sprintf("number: %d\nMessage: %+v\nReactions: %v", i, message.Msg, message.Msg.Reactions)
		fmt.Println(text)
		if message.Msg.ReplyCount > 0 {
			rparam := &slack.GetConversationRepliesParameters{
				ChannelID: CHANNEL,
				Timestamp: message.Msg.ThreadTimestamp,
				Cursor:    "",
				Inclusive: false,
				Latest:    message.Msg.LatestReply,
				Limit:     5,
				Oldest:    "",
			}
			rres, _, _, err := c.GetConversationReplies(rparam)
			if err != nil {
				panic(err)
			}
			fmt.Printf("-------リプライ----------")
			for j, rep := range rres {
				rtext := fmt.Sprintf("\n\t %d-%d\n\t%s", i, j, rep.Msg.Text)
				fmt.Println(rtext)
			}
		}
	}
}
