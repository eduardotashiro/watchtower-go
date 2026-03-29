package main

import (
	"fmt"

	// "github.com/eduardotashiro/watchtower-go/scraping"
	"github.com/eduardotashiro/watchtower-go/slack"

)

func main() {
	// app := scraping.CheckServiceStatus()
	bot :=slack.Bot()

	fmt.Println(bot)
}
