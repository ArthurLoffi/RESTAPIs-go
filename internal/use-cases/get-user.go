package use_cases

import (
	"net/http"
	user "restapis-go/internal/entities"

	"github.com/gin-gonic/gin"
)

// PostUser godoc
// @Summary      Adicionar um novo usuário
// @Description  Adiciona um novo usuário ao json
// @Tags         users
// @Produce      json
// @Success      201  {array}  user.User
// @Router       /users/post [post]
func NewUser(c *gin.Context) {
	var newUser user.User

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.Users = append(user.Users, newUser)

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created sucessfully",
		"data":    newUser,
	})
}
