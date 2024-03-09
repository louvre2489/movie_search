package main

import (
	"github.com/gin-gonic/gin"
	"github.com/louvre2489/movie_search/controller"
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
