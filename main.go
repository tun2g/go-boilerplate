package main

import (
	"os"

	server "fist-app/src"
	"fist-app/src/database"
	"fist-app/src/lib/logger"

	"github.com/urfave/cli"

	"github.com/joho/godotenv"
)

var (
	client *cli.App
)

func init() {
	client = cli.NewApp()
	client.Name = ""
	client.Usage = ""
	client.Version = "0.0.0"
}

func main() {
	var _logger = logger.NewLogger("main")

	if err := godotenv.Load(); err != nil {
		_logger.Fatal("Error loading .env file")
	}

	client.Commands = []cli.Command{
		// RUN: server
		server.StartServer(),

		// RUN: migrate
		database.Migration(),

		// RUN: rollback
		database.Rollback(),		
	}

	// Run the CLI app
	err := client.Run(os.Args)
	if err != nil {
		_logger.Fatalf(err.Error())
	}
}
