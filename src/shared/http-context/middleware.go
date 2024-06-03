package httpContext

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HttpContextMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		customContext := &CustomContext{Context: ctx}
		requestId := uuid.New()

		customContext.SetRequestId(requestId.String())

		ctx.Next()
	}
}
