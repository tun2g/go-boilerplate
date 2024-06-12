package exception

import (
	"fmt"
	"net/http"
	
	httpContext "fist-app/src/shared/http-context"
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

func ThrowUnprocessableEntityException(ctx *httpContext.CustomContext, err error){
	details := manufactureValidationException(err)
	ctx.AbortWithStatusJSON(
		http.StatusUnprocessableEntity,
		HttpError{
			Message: "Unprocessable Entity Exception",
			RequestId: ctx.GetRequestId(),
			Details: details,
		},
	)
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
