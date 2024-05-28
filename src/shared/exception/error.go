package exception

import (
	httpContext "fist-app/src/shared/http-context"
)

type ErrorDetail struct {
	Issue   string `json:"issue"`
	IssueId string `json:"issueId"`
	Field   string `json:"field"`
}

type HttpError struct {
	RequestId string        `json:"requestId"`
	Message   string        `json:"message"`
	Details   []ErrorDetail `json:"details"`
}

func (e *HttpError) Error() string {
	return e.Message
}

func NewHttpError(ctx *httpContext.CustomContext, requestId string, status int, message string, details []ErrorDetail) {
	ctx.AbortWithStatusJSON(status, HttpError{
		RequestId: requestId,
		Message:   message,
		Details:   details,
	})
}
