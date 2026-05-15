package main

import 	"github.com/gin-gonic/gin"

func setupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", getUsers)
		v1.POST("/users/post", postUser)
	}
}