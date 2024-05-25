package exception

func NewUnprocessableEntityException(requestId string, details []ErrorDetail) *HttpError {
	return &HttpError{
		RequestId: requestId,
		Status:    422,
		Message:   "Unprocessable Entity Exception",
		Details:   details,
	}
}
