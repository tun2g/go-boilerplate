package main

import (
	"log"
	"os"

	server "fist-app/src"
	"fist-app/src/config"
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


func startServer(){
	connection := database.InitDB()

	server, err := server.NewServer(connection, config.AppConfiguration)

	if(err != nil) {
		log.Print("Can not start server due to", err)
	}

	if err := server.Run(os.Getenv("APP_PORT")); err != nil {
		log.Print(err)
	}
}

func main() {
	var _logger = logger.Logger()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	client.Commands = []cli.Command{
		{
			Name:  "server",
			Usage: "send example tasks ",
			Action: func(c *cli.Context) error {
				startServer()
				return nil
			},
		},
		database.Migration(),
		database.Rollback(),		
	}

	// Run the CLI app
	err := client.Run(os.Args)
	if err != nil {
		_logger.Fatalf(err.Error())
	}

}
