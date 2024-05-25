package main

import (
	"log"
	"os"

	server "fist-app/src"
	"fist-app/src/db"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connection := db.InitDB()

	defer func() {
		if err := connection.DB().Close(); err != nil {
			log.Print(err)
		}
	}()

	server, err := server.NewServer(connection)

	if(err != nil) {
		log.Print("Can not start server due to", err)
	}

	if err := server.Run(os.Getenv("PORT")); err != nil {
		log.Print(err)
	}
}
