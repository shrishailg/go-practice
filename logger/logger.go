package logger

import (
	"io"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-practice/config"
	"github.com/sirupsen/logrus"
)

const XRequestId = "X-Request-Id"

func Info(ctx *gin.Context, message string) {
	logrus.WithFields(logrus.Fields{
		XRequestId: ctx.Request.Header.Get(XRequestId),
		"Path":     ctx.Request.RequestURI,
	}).Infoln(message)
}

func Debug(ctx *gin.Context, message string, err error) {
	logrus.WithFields(logrus.Fields{
		XRequestId: ctx.Request.Header.Get(XRequestId),
		"Path":     ctx.Request.RequestURI,
		"Error":    err.Error(),
	}).Debugln(message)
}

func Error(ctx *gin.Context, message string, err error) {
	logrus.WithFields(logrus.Fields{
		XRequestId: ctx.Request.Header.Get(XRequestId),
		"Path":     ctx.Request.RequestURI,
		"Error":    err.Error(),
	}).Errorln(message)
}

func Panic(message string) {
	logrus.Panicln(message)
}

func Fatal(message string) {
	logrus.Fatalln(message)
}

func Init() {
	// to add logging to some log file instead of console
	f, err := os.Create(config.Config.GetString("logging.path"))

	if err != nil {
		Fatal("error occured while creating file" + err.Error())
	}
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	//to set log levels
	logrus.SetLevel(getLogLevel())
	// logrus.SetLevel(logrus.InfoLevel)

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableTimestamp: true,
		FullTimestamp:    false,
	})

}

func getLogLevel() logrus.Level {

	level := config.Config.GetString("logging.level")

	switch strings.ToLower(level) {
	case "debug":
		return logrus.DebugLevel
	case "error":
		return logrus.ErrorLevel
	default:
		return logrus.InfoLevel
	}
}
