package exception

import (
	"fmt"

	validator "github.com/go-playground/validator/v10"
)

type UnprocessableEntityException struct {
	HttpError
}

func manufactureValidationException(err error) []ErrorDetail{
	var validationErrors []ErrorDetail

	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, e := range errs {
				validationErrors = append(validationErrors, ErrorDetail{
					Field: e.Field(),
					IssueId: e.Tag(),
					Issue: fmt.Sprintf("Validation failed on %s.", e.Field(),),
				})
			}
		} else {
			validationErrors = append(validationErrors, ErrorDetail{
				Field: "",
				Issue: err.Error(),
			})
		}
	}
	return validationErrors
}

func NewUnprocessableEntityException(requestId string, err error) *UnprocessableEntityException{
	details := manufactureValidationException(err)
	return &UnprocessableEntityException{
		HttpError: HttpError{
			RequestId: requestId,
			Message: "Unprocessable Entity Exception",
			Details: details,
		},
	}
}
