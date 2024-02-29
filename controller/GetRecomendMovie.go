package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetRecomendMovie(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"recomend_movie": "this one",
	})
}
