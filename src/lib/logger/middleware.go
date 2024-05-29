package logger

import (
	httpContext "fist-app/src/shared/http-context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

func RequestLoggerMiddleware(ctx *httpContext.CustomContext) {
	logger := Logger()
	startTime := time.Now()

	endTime := time.Now()
	latency := endTime.Sub(startTime)

	clientIP := ctx.ClientIP()
	method := ctx.Request.Method
	path := ctx.Request.URL.Path
	userAgent := ctx.Request.UserAgent()
	dataLength := ctx.Writer.Size()

	entry := logger.WithFields(logrus.Fields{
		"latency":     latency,
		"client_ip":   clientIP,
		"method":      method,
		"path":        path,
		"user_agent":  userAgent,
		"data_length": dataLength,
	})

	if len(ctx.Errors) > 0 {
		entry.Error(ctx.Errors.String())
	}
	entry.Info(fmt.Sprintf("----------Request received: %s", ctx.GetRequestId()))
}

func ResponseLoggerMiddleware(ctx *httpContext.CustomContext) {
	logger := Logger()

	ctx.Next()

	statusCode := ctx.Writer.Status()
	clientIP := ctx.ClientIP()
	method := ctx.Request.Method
	path := ctx.Request.URL.Path
	userAgent := ctx.Request.UserAgent()
	dataLength := ctx.Writer.Size()

	entry := logger.WithFields(logrus.Fields{
		"status_code": statusCode,
		"client_ip":   clientIP,
		"method":      method,
		"path":        path,
		"user_agent":  userAgent,
		"data_length": dataLength,
	})

	if len(ctx.Errors) > 0 {
		entry.Errorf("----------Request completed: %s due to %s", ctx.GetRequestId(), ctx.Errors.String())
	} else{
		entry.Info(fmt.Sprintf("----------Request completed: %s", ctx.GetRequestId()))
	}
}
