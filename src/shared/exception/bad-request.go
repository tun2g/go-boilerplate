package exception

func NewBadRequestException(requestId string, details []ErrorDetail) *HttpError {
	return &HttpError{
		RequestId: requestId,
		Status:    400,
		Message:   "Bad Request Exception",
		Details:   details,
	}
}