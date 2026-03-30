package slack

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/eduardotashiro/watchtower-go/scraping"
	"github.com/joho/godotenv"
)

// ligando watchtower
func Bot() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatal(err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	sb := string(body)
	fmt.Printf("%s", sb)

}

// Função que envia msg para o slack
func PostMessageSlack() {
	data := scraping.CheckServiceStatus()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Post(os.Getenv("IW"), "application/json", data)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	sb := string(body)
	log.Printf("status: %s", sb)
}
