package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    docs "git-search-api/cmd/docs"
)

// @title           Git Search API
// @version         1.0
// @description     API de exemplo
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
    
    // Serve Swagger UI at /swagger/index.html
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	setupRoutes(router)

	router.Run(":8080")
}