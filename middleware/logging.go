package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

func HttpLogToFile(appMode string)  {
	t:=fmt.Sprintf("%s",time.Now())[:10]
	fileName := "./runtime/http_logs/"+t+"-gin_http.log"		//写入文件
	logfile, err := os.OpenFile(fileName,os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("err", err)
	}
	gin.SetMode(appMode)
	gin.DefaultWriter = io.MultiWriter(logfile)
}

// 日志记录到文件
func LoggerToFile() gin.HandlerFunc {
	t := fmt.Sprintf("%s", time.Now())[:10]
	fileName := "./runtime/logs/" + t + "-info.log" //写入文件
	src, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0766)
	if err != nil {
		fmt.Println("err", err)
	}
	logger := logrus.New()	//实例化
	logger.Out = src		//设置输出
	logger.SetLevel(logrus.DebugLevel)			//设置日志级别
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",	//设置日志格式
	})
	return func(c *gin.Context) {
		startTime := time.Now()		// 开始时间
		c.Next()					// 处理请求
		endTime := time.Now()		// 结束时间
		latencyTime := endTime.Sub(startTime)		// 执行时间
		reqMethod := c.Request.Method				// 请求方式
		reqUri := c.Request.RequestURI				// 请求路由
		statusCode := c.Writer.Status()				// 状态码
		clientIP := c.Request.Host					// 请求IP
		logger.WithFields(logrus.Fields{			// 日志格式
			"statusCode":  statusCode,
			"latencyTime": latencyTime,
			"clientIp":    clientIP,
			"reqMethod":   reqMethod,
			"reqUri":      reqUri,
		}).Info()
	}
}