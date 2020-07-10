//middlewares for project
package middleware

import (
	"fmt"
	"os"
	"time"

	"git.ont.io/ontid/otf/utils"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var Log *logrus.Logger

func LoggerToFile() gin.HandlerFunc {
	logFilePath := utils.DEFAULT_LOG_FILE_PATH
	if fi, err := os.Stat(logFilePath); err == nil {
		if !fi.IsDir() {
			fmt.Printf("open %s: not a directory", logFilePath)
			panic(err)
		}
	} else if os.IsNotExist(err) {
		if err := os.MkdirAll(logFilePath, 0766); err != nil {
			fmt.Println("err", err)
			panic(err)
		}
	} else {
		fmt.Println("err", err)
		panic(err)
	}
	nowTime := time.Now().Format("2006-01-02_15.04.05")
	src, err := os.OpenFile(logFilePath+nowTime+"_Log.log", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	Log = logrus.New()
	Log.Out = src
	Log.SetLevel(logrus.DebugLevel)
	Log.SetOutput(os.Stdout)
	Log.SetReportCaller(true)
	Log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logWriter, err := rotatelogs.New(
		logFilePath+nowTime+"_Log.log",
		rotatelogs.WithLinkName(logFilePath),
		rotatelogs.WithMaxAge(7*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	if err != nil {
		panic(err)
	}
	logrus.SetOutput(logWriter)
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		Log.WithFields(logrus.Fields{
			"status_code":  statusCode,
			"latency_time": latencyTime,
			"client_ip":    clientIP,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
		}).Info()
	}
}
