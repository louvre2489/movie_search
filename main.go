package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
  engine := gin.Default()
  engine.GET("/", getRoute)

  engine.Run()
}

func getRoute(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{
    "message": "hello world!",
  })
}
