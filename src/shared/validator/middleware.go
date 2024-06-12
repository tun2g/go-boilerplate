package validator

import (
	httpContext "fist-app/src/shared/http-context"
	"fist-app/src/shared/exception"
	
)

func BindJson[T any](dto *T) func(ctx *httpContext.CustomContext) {
	return func(ctx *httpContext.CustomContext) {
		if err := ctx.ShouldBindJSON(dto); err != nil {
			if err := ctx.ShouldBindJSON(&dto); err != nil {
				print(err)
				exception.ThrowUnprocessableEntityException(ctx, err);
				return
			}				
		}
		ctx.Next()
	}
}