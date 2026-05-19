package main

import (
	docs "git-search-api/docs"
	"git-search-api/internal/repository"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           Git Search API
// @version         1.0
// @description     API de exemplo
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	router := gin.Default()

	repository.Connect()

	docs.SwaggerInfo.BasePath = "/api/v1"
    
    // Serve Swagger UI at /swagger/index.html
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/healty", func (c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})
	setupRoutes(router)

	router.Run(":8080")
}