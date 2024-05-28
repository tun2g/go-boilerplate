package httpContext

import "github.com/gin-gonic/gin"

func CustomContextHandler(handler func(ctx *CustomContext)) gin.HandlerFunc {
	return func(_ctx *gin.Context) {
		customContext := &CustomContext{Context: _ctx}
		handler(customContext)
	}
}