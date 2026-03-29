package main

import (
	"fmt"

	"github.com/eduardotashiro/watchtower-go/scraping"
)

func main() {
	app := scraping.CheckServiceStatus()
	fmt.Println(app)	
}

