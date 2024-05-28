package exception

type BadRequestException struct {
	HttpError
}

func NewBadRequestException(requestId string, details []ErrorDetail) *BadRequestException {
	return &BadRequestException{
		HttpError: HttpError{
			RequestId: requestId,
			Message:   "Bad Request Exception",
			Details:   details,
		},
	}
}
