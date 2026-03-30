package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

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
	err := godotenv.Load()
	if err!= nil {
		log.Fatal(err)
	}

	postBody, _ := json.Marshal(map[string]string{
		"text": "Hello, World!",
	})
	responseBody := bytes.NewBuffer(postBody)
	resp, err := http.Post(os.Getenv("IW"), "application/json", responseBody)
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
