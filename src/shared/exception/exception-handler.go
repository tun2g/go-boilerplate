package exception

import (
	"net/http"

	"github.com/gin-gonic/gin"
	
	"fist-app/src/lib/logger"
	httpContext "fist-app/src/shared/http-context"
)

var _logger = logger.NewLogger("exception")

func ErrorHandler(ctx *httpContext.CustomContext) {
	ctx.Next()

	if len(ctx.Errors) > 0 {
		err := ctx.Errors[0]

		switch e := err.Err.(type) {
		case *UnauthorizedException:
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, e)
			return
		case *BadRequestException:
			ctx.AbortWithStatusJSON(http.StatusBadRequest, e)
			return
		case *UnprocessableEntityException:
			ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, e)
			return
		default:
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"requestId": ctx.GetRequestId(),
				"message":   "Internal Server Error Exception",
				"details":   []ErrorDetail{},
			})
			return
		}
	}
}
