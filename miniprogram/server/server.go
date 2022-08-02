package server

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

type ServerInterface interface {
	Run(add ...string) error
}

func NewServer(logPath string) ServerInterface {
	engine := gin.Default()

	setRoute(engine)
	setLogger(engine, logPath)
	return engine
}

func setLogger(engine *gin.Engine, logPath string) {
	f, err := os.OpenFile(logPath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("err %v", err)
	}

	logger := logrus.New()
	logger.Out = f
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{})

	loggerToFile := func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		logger.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}

	engine.Use(loggerToFile)
}
