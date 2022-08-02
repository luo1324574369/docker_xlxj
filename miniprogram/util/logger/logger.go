package logger

import (
	"github.com/sirupsen/logrus"
	"log"
	"miniprogram/util/config"
	"os"
)

type logger struct {
	writer *logrus.Logger
}

func NewLogger() *logger {
	logPath := config.RootPath + "/" + config.GetString("log.name")

	f, err := os.OpenFile(logPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0)
	if err != nil {
		log.Fatalf("err %v", err)
	}

	l := logrus.New()
	l.Out = f
	l.SetLevel(logrus.DebugLevel)
	l.SetFormatter(&logrus.TextFormatter{})

	return &logger{writer: l}
}

func (l *logger) Info(content string) {
	logrus.Info(content)
	l.writer.Info(content)
}

func (l *logger) Error(content string) {
	logrus.Error(content)
	l.writer.Error(content)
}
