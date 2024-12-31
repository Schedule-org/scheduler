package logging

import (
	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func InitLogger() *logrus.Logger {
	Log = logrus.New()
	Log.SetFormatter(&logrus.JSONFormatter{})
	Log.SetLevel(logrus.InfoLevel)

	logrus.SetOutput(&lumberjack.Logger{
		Filename:   "/var/log/app.log",
		MaxSize:    10,
		MaxBackups: 3,
		MaxAge:     28,
		Compress:   true,
	})

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	return Log
}
