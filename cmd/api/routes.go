package main

import (
	controller "restapis-go/cmd/api/controllers"
	"restapis-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.Use(middleware.Logger())

	{
		v1.GET("/users", controller.ListUsers)
		v1.GET("/users/:ID", controller.ListUserByID)
		v1.POST("/users/:name", controller.NewUser)
		v1.DELETE("/users/delete/:ID", controller.DeleteUser)
		v1.PATCH("/users/update/:ID/:name", controller.UpdateUser)
	}
}
