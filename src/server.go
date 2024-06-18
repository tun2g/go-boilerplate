package src

import (
	"context"
	"fist-app/src/config"
	"fist-app/src/database"
	logger "fist-app/src/lib/logger"
	"fist-app/src/shared/cors"
	"fist-app/src/shared/exception"
	httpContext "fist-app/src/shared/http-context"
	"os"

	_ "fist-app/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/urfave/cli"
	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	route  *gin.Engine
	ctx    context.Context
	config config.Config
}

var _logger = logger.NewLogger("server")

func NewServer(dbConnection *gorm.DB, config config.Config) (*Server, error) {
	ctx := context.Background()

	route := gin.Default()
	
	gin.DefaultWriter = _logger.Writer()
	
	route.Use(cors.CorsMiddleware())
	
	route.Use(httpContext.HttpContextMiddleware())

	route.Use(httpContext.CustomContextHandler(logger.RequestLoggerMiddleware))
	route.Use(httpContext.CustomContextHandler(logger.ResponseLoggerMiddleware))
	route.Use(httpContext.CustomContextHandler(exception.ErrorHandler))
	
	if(config.SwaggerEnabled){
		route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	server := &Server{
		db:     dbConnection,
		ctx:    ctx,
		route:  route,
		config: config,
	}

	server.launchingServer(route)
	return server, nil
}

func (server *Server) Run(addr string) error {
	return server.route.Run(":" + addr)
}

func (server *Server) Engine() *gin.Engine {
	return server.route
}

func (server *Server) Database() *gorm.DB {
	return server.db
}


func StartServer() cli.Command{
	cli := cli.Command{
		Name:  "server",
		Usage: "send example tasks ",
		Action: func(c *cli.Context) error {
			connection := database.InitDB()

			server, err := NewServer(connection, config.AppConfiguration)
		
			if(err != nil) {
				_logger.Print("Can not start server due to", err)
			}
		
			if err := server.Run(os.Getenv("APP_PORT")); err != nil {
				_logger.Print(err)
			}
			return nil
		},
	}
	return cli
}