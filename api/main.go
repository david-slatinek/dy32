package main

import (
	"log"
	"main/env"
	"os"
)

func main() {
	err := env.Load("env/.env")
	if err != nil {
		log.Fatalf("error with env.Load: %v", err)
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("MONGODB_URI is not set")
	}
}
