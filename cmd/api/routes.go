package main

import (
	controller "restapis-go/cmd/api/controllers"
	"restapis-go/internal/middleware"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerfiles "github.com/swaggo/files"
)

func setupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	v1.Use(middleware.Limiter())
	
	v1.POST("/login", controller.Login)

	v1.GET("/healthy", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})
	
	protect := v1.Group("/")
	protect.Use(middleware.Auth())

	{
		protect.GET("/users", controller.ListUsers)
		protect.GET("/user/:ID", controller.ListUserByID)
		protect.POST("/post", controller.NewUser)
		protect.DELETE("/delete/:ID", controller.DeleteUser)
		protect.PATCH("/update/:ID", controller.UpdateUser)
	}

	// Serve Swagger UI at /swagger/index.html
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
