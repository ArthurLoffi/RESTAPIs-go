package main

import (
	"fmt"
	docs "restapis-go/docs"
	"restapis-go/internal/repository"
	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           REST API go
// @version         1.0
// @description     API de exemplo
// @host            localhost:8080
// @BasePath        /api/v1
func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "DELETE", "PATCH"},
		AllowHeaders: []string{"Content-Type"},
	}))
	
	repository.Connect()
	
	docs.SwaggerInfo.BasePath = "/api/v1"
	
	printBanner()
	// Serve Swagger UI at /swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	router.GET("/healty", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"success": true,
		})
	})
	setupRoutes(router)

	router.Run(":8080")
}

func printBanner() {
	fmt.Println(`
╔══════════════════════════════════════════════════════════════╗
║                                                              ║
║   ██████╗ ███████╗███████╗████████╗     █████╗ ██████╗ ██╗   ║
║   ██╔══██╗██╔════╝██╔════╝╚══██╔══╝    ██╔══██╗██╔══██╗██║   ║
║   ██████╔╝█████╗  ███████╗   ██║       ███████║██████╔╝██║   ║
║   ██╔══██╗██╔══╝  ╚════██║   ██║       ██╔══██║██╔═══╝ ██║   ║
║   ██║  ██║███████╗███████║   ██║       ██║  ██║██║     ██║   ║
║   ╚═╝  ╚═╝╚══════╝╚══════╝   ╚═╝       ╚═╝  ╚═╝╚═╝     ╚═╝   ║
║                                                              ║
╚══════════════════════════════════════════════════════════════╝
`)
}