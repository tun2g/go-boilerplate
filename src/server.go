package src

import (
	"context"
	"fist-app/src/config"
	"fist-app/src/shared/cors"
	"fist-app/src/shared/exception"
	httpContext "fist-app/src/shared/http-context"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Server struct {
	db     *gorm.DB
	route  *gin.Engine
	ctx    context.Context
	config config.Config
}

func NewServer(dbConnection *gorm.DB, config config.Config) (*Server, error) {
	ctx := context.Background()

	route := gin.Default()
	
	route.Use(cors.CorsMiddleware())
	
	route.Use(httpContext.HttpContextMiddleware())
	
	route.Use(httpContext.CustomContextHandler(exception.ErrorHandler))

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
