package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	swaggerfiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    docs "git-search-api/cmd/docs"
)

type user struct {
	ID string `json:"ID"`
	Name string `json:"Nome"`
}

var users = []user{
	{ID: "1", Name: "Arthur"},
	{ID: "2", Name: "Caio"},
}

// GetUsers godoc
// @Summary      Lista todos os usuários
// @Description  Retorna a lista de usuários
// @Tags         users
// @Produce      json
// @Success      200  {array}  user
// @Router       /users [get]
func getUsers(c *gin.Context)  {
	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
    
    // Serve Swagger UI at /swagger/index.html
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	setupRoutes(router)

	router.Run(":8080")
}