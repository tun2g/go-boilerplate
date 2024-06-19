package utils

import (
	"errors"
	"fist-app/src/lib/logger"
	"fist-app/src/shared/exception"
	httpContext "fist-app/src/shared/http-context"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var _logger = logger.NewLogger("utils")

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

func UUIDParamsMiddleware(fields ...string) func(ctx *httpContext.CustomContext) {
	if len(fields) == 0 {
		fields = []string{"id"}
	}

	return func(ctx *httpContext.CustomContext) {
		var wg sync.WaitGroup
		errCh := make(chan error, 1) // Buffer of 1 to ensure we can send an error without blocking

		for _, field := range fields {
			wg.Add(1)
			go func(field string) {
				defer wg.Done()
				id := ctx.Param(field)
				if _, err := uuid.Parse(id); err != nil {
					select {
					case errCh <- errors.New("Invalid UUID Params for field: " + field):
					default:
					}
				}
			}(field)
		}

		wg.Wait()
		close(errCh)

		if err := <-errCh; err != nil {
			exception.ThrowUnprocessableEntityException(ctx, err)
			return
		}

		ctx.Next()
	}
}
