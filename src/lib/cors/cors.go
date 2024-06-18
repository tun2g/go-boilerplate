package cors

import (
	"github.com/gin-contrib/cors"
)

var CorsConfig = cors.New(cors.Config{
	AllowOrigins:     []string{"*"},
	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
	AllowCredentials: true,
	AllowHeaders:     []string{"Content-Type", "Content-Length", "api_key", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
})
