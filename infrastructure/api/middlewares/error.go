package middlewares

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/kmaguswira/coinbit/utils/error_mapping"
)

type newErrorMiddleware struct{}

func NewErrorMiddleware() *newErrorMiddleware {
	return &newErrorMiddleware{}
}

func (t *newErrorMiddleware) handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		lastErr := c.Errors.Last()
		if lastErr == nil {
			return
		}

		for _, e := range error_mapping.Errors {
			for _, e2 := range e.FromErrors {
				if errors.Is(lastErr.Err, e2) {
					e.Response(c, lastErr.Err)
					return
				}
			}
		}
	}
}

func (t *newErrorMiddleware) GetHandler() gin.HandlerFunc {
	return t.handler()
}
