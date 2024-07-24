package app

import (
	"os"

	"github.com/ilovepitsa/youtube_tumbnail/config"

	"github.com/sirupsen/logrus"
)

func SetLogrus(level config.Logger) {
	logrusLevel, err := logrus.ParseLevel(level.Level)
	if err != nil {
		logrus.SetLevel(logrus.DebugLevel)
	} else {
		logrus.SetLevel(logrusLevel)
	}

	if level.IsJSON {
		logrus.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	} else {
		logrus.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}

	logrus.SetOutput(os.Stdout)
}
