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
	Name string `json:"Name"`
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

// PostUser godoc
// @Summary      Adicionar um novo usuário
// @Description  Adiciona um novo usuário ao json
// @Tags         users
// @Produce      json
// @Success      201  {array}  user
// @Router       /users/post [post]
func postUser(c *gin.Context) {
	var newUser user

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users = append(users, newUser)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"data": newUser,
	})
}

func main() {
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api/v1"
    
    // Serve Swagger UI at /swagger/index.html
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	setupRoutes(router)

	router.Run(":8080")
}