package middlewares

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/coinbit/utils/logger"
)

type newLoggerMiddleware struct {
	logger logger.LoggerStruct
}

func NewLoggerMiddleware() *newLoggerMiddleware {
	return &newLoggerMiddleware{
		logger: logger.Log(),
	}
}

func (t *newLoggerMiddleware) handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		t.logger.Log.Infof("| %3d | %13v | %15s | %s | %s |",
			statusCode,
			latencyTime,
			clientIP,
			reqMethod,
			reqUri,
		)
	}

}

func (t *newLoggerMiddleware) GetHandler() gin.HandlerFunc {
	return t.handler()
}
