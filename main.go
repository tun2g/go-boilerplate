package main

import (
	"log"
	"os"

	server "fist-app/src"
	"fist-app/src/config"
	"fist-app/src/database"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connection := database.InitDB()

	server, err := server.NewServer(connection, config.AppConfiguration)

	if(err != nil) {
		log.Print("Can not start server due to", err)
	}

	if err := server.Run(os.Getenv("PORT")); err != nil {
		log.Print(err)
	}
}
