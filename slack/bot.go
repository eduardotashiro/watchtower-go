package slack

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
	"github.com/slack-go/slack/socketmode"
)

// ligando watchtower
func Bot() *socketmode.Client {
	sbt := os.Getenv("SBT")
	sat := os.Getenv("SAT")

	// fmt.Println(sat, sbt)
	api := slack.New(
		sbt,
		slack.OptionDebug(true),
	)

	client := socketmode.New(
		api,
		socketmode.OptionDebug(true),
	)

	return client
}
