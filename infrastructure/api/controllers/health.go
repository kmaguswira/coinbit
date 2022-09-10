package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthController struct{}

func NewHealthController() *healthController {
	return &healthController{}
}

func (t healthController) Check(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
