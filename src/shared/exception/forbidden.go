package exception

func NewForbiddenException(requestId string) *HttpError {
	return &HttpError{
		RequestId: requestId,
		Status:    403,
		Message:   "Forbidden Exception",
		Details:   []ErrorDetail{},
	}
}