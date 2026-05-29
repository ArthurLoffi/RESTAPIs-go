package main

import (
	controller "restapis-go/cmd/api/controllers"
	"restapis-go/internal/middleware"

	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	
	v1.POST("/login", controller.Login)
	
	protect := v1.Group("/")
	protect.Use(middleware.Auth())
	{
		protect.GET("/", controller.ListUsers)
		protect.GET("/:ID", controller.ListUserByID)
		protect.POST("/post", controller.NewUser)
		protect.DELETE("/delete/:ID", controller.DeleteUser)
		protect.PATCH("/update/:ID/:name", controller.UpdateUser)
	}
}
