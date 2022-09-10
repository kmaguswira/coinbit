package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/coinbit/infrastructure/external_services/redis"
	"github.com/kmaguswira/coinbit/utils/logger"
)

type idempotentMiddleware struct{}

type Idempotent struct {
	RequestID    string `json:"requestId"`
	HttpStatus   int    `json:"httpStatus"`
	Url          string `json:"url"`
	Method       string `json:"method"`
	RequestBody  string `json:"requestBody"`
	ResponseBody string `json:"responseBody"`
}

type responseBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func (r responseBodyWriter) WriteString(s string) (n int, err error) {
	r.body.WriteString(s)
	return r.ResponseWriter.WriteString(s)
}

func NewIdempotentMiddleware() *idempotentMiddleware {
	return &idempotentMiddleware{}
}

func (t *idempotentMiddleware) handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		rid := c.Request.Header.Get("X-Request-ID")

		if rid == "" {
			c.Next()
			return
		}

		val, _ := redis.RedisClient.GetValue(fmt.Sprintf("idempotent_request_api_%s", rid))

		if val != "" {
			idempotent := Idempotent{}
			json.Unmarshal([]byte(val), &idempotent)

			if idempotent.ResponseBody == "" {
				c.AbortWithStatus(idempotent.HttpStatus)
				return
			}

			responseBody := make(map[string]interface{})
			if err := json.Unmarshal([]byte(idempotent.ResponseBody), &responseBody); err != nil {
				logger.Log().Error(err)
			}

			c.AbortWithStatusJSON(idempotent.HttpStatus, responseBody)
			return
		}

		w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: c.Writer}
		c.Writer = w

		c.Next()

		result := Idempotent{
			RequestID:    c.Request.Header.Get("X-Request-ID"),
			HttpStatus:   c.Writer.Status(),
			Url:          c.Request.URL.RequestURI(),
			Method:       c.Request.Method,
			ResponseBody: w.body.String(),
		}

		body, isExist := c.Get("body")
		if isExist {
			result.RequestBody = body.(string)
		}

		idempotentString, err := json.Marshal(result)
		if err != nil {
			logger.Log().Error(err)
			return
		}
		redis.RedisClient.SetValue(fmt.Sprintf("idempotent_request_api_%s", rid), string(idempotentString), 1*60)
	}
}

func (t *idempotentMiddleware) GetHandler() gin.HandlerFunc {
	return t.handler()
}
