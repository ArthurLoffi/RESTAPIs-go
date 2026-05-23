package main

import (
	controller "restapis-go/cmd/api/controllers"

	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	{
		v1.GET("/users", controller.ListUsers)
		v1.POST("/users/:name", controller.NewUser)
		v1.DELETE("/users/delete/:ID", controller.DeleteUser)
	}
}
