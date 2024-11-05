package main

import (
	"github.com/gin-gonic/gin"
	"test_server/controller"
)

func main() {
	router := gin.Default()
	router.Use(gin.Recovery())
	v1 := router.Group("/v1")
	{
		v1.GET("/get", controller.List)
	}
	router.Run("127.0.0.1:8000")
}
