package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	s3Bucket := os.Getenv("S3_BUCKET")
	fmt.Println(s3Bucket)

	//POC
	//BC
	//"window.PogoConfig = {\"template\":\"status\",\"service\":\"pokerstars\",\"category\":\"Betting, Lotteries & Fantasy Sports\",\"outage\":false};",
}
