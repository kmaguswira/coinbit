package middlewares

import (
	"github.com/ahmetb/go-linq"

	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/coinbit/application/config"
)

type newCorsMiddleware struct {
}

func NewCorsMiddleware() *newCorsMiddleware {
	return &newCorsMiddleware{}
}

func (t *newCorsMiddleware) handler() gin.HandlerFunc {
	return func(c *gin.Context) {

		allowedOrigin := config.GetConfig().CorsAllowOrigin

		requestHeader := "*"

		if len(c.Request.Header["Origin"]) > 0 {
			requestHeader = c.Request.Header["Origin"][0]
		}

		if linq.From(allowedOrigin).Contains(requestHeader) {
			c.Writer.Header().Set("Access-Control-Allow-Origin", requestHeader)
		}

		c.Writer.Header().Set("Access-Control-Max-Age", config.GetConfig().CorsMaxAge)
		c.Writer.Header().Set("Access-Control-Allow-Methods", config.GetConfig().CorsAllowMethods)
		c.Writer.Header().Set("Access-Control-Allow-Headers", config.GetConfig().CorsAllowHeaders)
		c.Writer.Header().Set("Access-Control-Expose-Headers", config.GetConfig().CorsAllowHeaders)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", config.GetConfig().CorsAllowCredentials)

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func (t *newCorsMiddleware) GetHandler() gin.HandlerFunc {
	return t.handler()
}
