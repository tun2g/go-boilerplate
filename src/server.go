package src

import (
	"context"
	"fist-app/src/shared/cors"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Server struct {
	engine *gin.Engine
	db     *gorm.DB
	route  *gin.Engine
	ctx    context.Context
}

func NewServer(dbConnection *gorm.DB) (*Server, error) {
	ctx := context.Background()

	route := gin.Default()

	route.Use(cors.CorsMiddleware())

	server := &Server{
		engine: gin.Default(),
		db:     dbConnection,
		ctx:    ctx,
		route:  route,
	}

	server.setupRoutes(route)	

	return server, nil
}

func (server *Server) Run(addr string) error {
	return server.engine.Run(":" + addr)
}

func (server *Server) Engine() *gin.Engine {
	return server.engine
}

func (server *Server) Database() *gorm.DB {
	return server.db
}
