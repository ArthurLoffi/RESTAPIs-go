package main

import (
	"git-search-api/entity/user"
	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	{
		v1.GET("/users", user.GetUsers)
		v1.POST("/users/post", user.NewUser)
		v1.DELETE("/users/delete/:ID", user.DeleteUser)
	}
}