package exception

import (
	httpContext "fist-app/src/shared/http-context"
	"fmt"
	"encoding/json"
)

type ErrorDetail struct {
	Issue   string `json:"issue,omitempty"`
	IssueId string `json:"issueId,omitempty"`
	Field   string `json:"field,omitempty"`
}

type HttpError struct {
	RequestId string        `json:"requestId"`
	Message   string        `json:"message"`
	Details   []ErrorDetail `json:"details"`
}

func (e *HttpError) Error() string {
	errMsg := fmt.Sprintf("RequestID: %s, Message: %s", e.RequestId, e.Message)

	if len(e.Details) > 0 {
		detailsJSON, _ := json.Marshal(e.Details)
		errMsg += fmt.Sprintf(", Details: %s", detailsJSON)
	}

	return errMsg
}

func NewHttpError(ctx *httpContext.CustomContext, requestId string, status int, message string, details []ErrorDetail) {
	ctx.AbortWithStatusJSON(status, HttpError{
		RequestId: requestId,
		Message:   message,
		Details:   details,
	})
}
