package exception

type ForbiddenException struct {
	HttpError
}

func NewForbiddenException(requestId string) *ForbiddenException{
	return &ForbiddenException{
		HttpError: HttpError{
			RequestId: requestId,
			Message: "Forbidden",
			Details: []ErrorDetail{{}},
		},
	}
}
