package main

import (
	"fist-app/src/shared/exception"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandlingMiddleware(c *gin.Context) {
	c.Next()

	httpError := exception.NewHttpError(
		c.GetString("requestId"),
		500,
		"Internal Server Error",
		[]exception.ErrorDetail{
			{Issue: "Internal Server Error", IssueId: "unknown"},
		},
	)
	c.JSON(http.StatusInternalServerError, httpError)
}
