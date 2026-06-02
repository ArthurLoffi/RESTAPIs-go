package main

import (
	"fmt"
	"os"
	docs "restapis-go/docs"
	"restapis-go/internal/middleware"
	"restapis-go/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// @title           REST API go
// @version         2.0
// @description     API de exemplo
// @host            localhost:8080
// @BasePath        /api/v1
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	godotenv.Load()
	ginMode := os.Getenv("GIN_MODE")
	if ginMode == "" {
		ginMode = "debug"
	}
	gin.SetMode(ginMode)

	router := gin.New()

	// Garantir que não caia se tiver um panic
	router.Use(gin.Recovery())
	router.Use(middleware.Logger())
	
	repository.Connect()
	
	docs.SwaggerInfo.BasePath = "/api/v1"
	
	printBanner()

	setupRoutes(router)

	router.Run(":8080")
}

func printBanner() {
	fmt.Print(`
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