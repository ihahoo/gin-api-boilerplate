package log

import (
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// GetLog ...
func GetLog() *logrus.Logger {
	log.Formatter = &logrus.JSONFormatter{}
	os.MkdirAll("data/logs/", 0777)
	file, err := os.OpenFile("data/log/app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		log.Out = file
	} else {
		log.Info("Failed to log to file, using default stderr")
	}
	return log
}
