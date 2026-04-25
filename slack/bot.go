package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/eduardotashiro/watchtower-go/scraping"
	"github.com/joho/godotenv"
)

// función que envía un mensaje a Slack
func PostMessageSlack() {
	payload := scraping.CheckServiceStatus()

	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("send?", string(jsonData))

	erro := godotenv.Load()
	if erro != nil {
		log.Fatal(erro)
	}

	resp, err := http.Post(os.Getenv("IW"), "application/json", bytes.NewReader(jsonData))
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if resp.StatusCode == 200 {
		fmt.Printf("Status Code: %d\n", resp.StatusCode)
		fmt.Printf("Response: %s", string(body))
	} else {
		fmt.Printf("error en slack, status %v", resp.StatusCode)
	}

}
