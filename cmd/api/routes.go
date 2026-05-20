package main

import (
	use_cases "restapis-go/internal/use-cases"

	"github.com/gin-gonic/gin"
)

func setupRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")

	{
		v1.GET("/users", use_cases.GetUsers)
		v1.POST("/users/post", use_cases.NewUser)
		v1.DELETE("/users/delete/:ID", use_cases.DeleteUser)
	}
}
