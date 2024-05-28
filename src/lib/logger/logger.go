package logger

import (
	"sync"

	"github.com/sirupsen/logrus"
)

var (
	loggerInstance *logrus.Logger
	once           sync.Once
)

// Logger returns a singleton instance of the logrus.Logger
func Logger() *logrus.Logger {
	once.Do(func() {
		loggerInstance = logrus.New()
		loggerInstance.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
			PadLevelText:  true,
		})
	})
	return loggerInstance
}
