package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRoute(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world!",
	})
}
