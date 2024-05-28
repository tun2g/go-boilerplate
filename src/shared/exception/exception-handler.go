package exception

import (
	"fist-app/src/lib/logger"
	httpContext "fist-app/src/shared/http-context"
	"net/http"

	"github.com/gin-gonic/gin"
)

var _logger = logger.Logger()

func ErrorHandler(ctx * httpContext.CustomContext){
		ctx.Next()

		_logger.Error(ctx.Errors)

		if len(ctx.Errors) > 0 {
			err := ctx.Errors[0].Err

			switch e := err.(type) {
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
					"details":   []ErrorDetail{{}},
				})
			}
		}
}
