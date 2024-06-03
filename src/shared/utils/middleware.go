package utils

import (
	"github.com/gin-gonic/gin"
)

func CombineMiddlewares(middlewares ...gin.HandlerFunc) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		for _, middleware := range middlewares {
			if middleware != nil {
				middleware(ctx)
				if ctx.IsAborted() {
					return
				}
			}
		}
	}
}