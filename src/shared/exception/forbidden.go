package exception

import (
	httpContext "fist-app/src/shared/http-context"

	"net/http"
)

type ForbiddenException struct {
	HttpError
}

func NewForbiddenException(requestId string) *ForbiddenException {
	return &ForbiddenException{
		HttpError: HttpError{
			RequestId: requestId,
			Message:   "Forbidden",
			Details:   []ErrorDetail{},
		},
	}
}

func ThrowForbiddenException(ctx *httpContext.CustomContext, details ...[]ErrorDetail) {
	var errorDetails []ErrorDetail
	if len(details) > 0 {
		errorDetails = details[0]
	} else {
		errorDetails = []ErrorDetail{}
	}

	ctx.AbortWithStatusJSON(
		http.StatusForbidden,
		HttpError{
			Message:   "Forbidden",
			RequestId: ctx.GetRequestId(),
			Details:   errorDetails,
		},
	)
}
