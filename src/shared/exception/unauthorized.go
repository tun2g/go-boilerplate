package exception

func NewUnauthorizedException(requestId string) *HttpError {
	return &HttpError{
		RequestId: requestId,
		Status:    401,
		Message:   "Unauthorized Exception",
		Details:   []ErrorDetail{},
	}
}