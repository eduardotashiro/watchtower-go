package main

import (
	"fmt"
	"time"

	"github.com/eduardotashiro/watchtower-go/slack"
)

// punto de entrada
func main() {
	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for t := range ticker.C {
		fmt.Println("tarea ejecutada en ", t)
		slack.PostMessageSlack()
	}

}
