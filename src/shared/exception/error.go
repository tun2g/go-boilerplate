package exception

type ErrorDetail struct {
	Issue   string `json:"issue"`
	IssueId string `json:"issueId"`
	Message string `json:"message"`
}

type HttpError struct {
	RequestId string        `json:"requestId"`
	Status    int           `json:"status"`
	Message   string        `json:"message"`
	Details   []ErrorDetail `json:"details"`
}

func NewHttpError(requestId string, status int, message string, details []ErrorDetail) *HttpError {
	return &HttpError{
		RequestId: requestId,
		Status:    status,
		Message:   message,
		Details:   details,
	}
}
