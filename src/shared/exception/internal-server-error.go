package exception

type InternalServerError struct {
	HttpError
}

func NewInternalServerError(requestId string) *InternalServerError{
	return &InternalServerError{
		HttpError: HttpError{
			RequestId: requestId,
			Message: "Internal Server Error Exception",
			Details: []ErrorDetail{},
		},
	}
}