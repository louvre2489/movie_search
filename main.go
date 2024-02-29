package main

import (
	"movie_search/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// v1 のグループ
	v1 := router.Group("/v1")
	{
		v1.GET("", controller.GetRoute)
		v1.GET("/", controller.GetRoute)
		v1.GET("/search", controller.GetRecomendMovie)
	}

	router.Run()
}
