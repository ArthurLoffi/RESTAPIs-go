package user

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type User struct {
	ID string `json:"ID"`
	Name string `json:"Name"`
}

var users = []User{}

// GetUsers godoc
// @Summary      Lista todos os usuários
// @Description  Retorna a lista de usuários
// @Tags         users
// @Produce      json
// @Success      200  {array}  User
// @Router       /users [get]
func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

// PostUser godoc
// @Summary      Adicionar um novo usuário
// @Description  Adiciona um novo usuário ao json
// @Tags         users
// @Produce      json
// @Success      201  {array}  User
// @Router       /users/post [post]
func NewUser(c *gin.Context) {
	var newUser User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	users = append(users, newUser)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created sucessfully",
		"data": newUser,
	})
}

// DeleteUser godoc
// @Summary      Remover usuário por id
// @Description  Remove o usuário com o determinado id no json
// @Tags         users
// @Produce      json
// @Success      200  {array}  User
// @Router       /users/delete [post]
func DeleteUser(c *gin.Context) {
	idUser := c.Param("ID")

	for i, u := range users {
		if u.ID == idUser {
			users = append(users[:i], users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Deleted user sucessfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}