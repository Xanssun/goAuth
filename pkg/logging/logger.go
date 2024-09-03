package logging

import (
	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func InitLogger() {
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	Log.SetLevel(logrus.InfoLevel)
}
