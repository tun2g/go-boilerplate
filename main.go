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


// @title Swagger UI
// @version 1.0
// @description Golang Gin Boilerplate.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
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
