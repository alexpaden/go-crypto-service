package main

import (
	"log"

	"github.com/alexpaden/go-crypto-service/pkg/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Panicf("Error loading main .env file")
	}
	server := server.NewServer()

	server.Run()

}
