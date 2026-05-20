package use_cases

import (
	"net/http"
	user "restapis-go/internal/entities"
	db "restapis-go/internal/repository"

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
	nameNewUser := c.Param("name")

	if nameNewUser == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Name can't empty",
		})
		return
	}

	newUser := user.User{
		Name: nameNewUser,
	}

	result := db.Database.Create(&newUser)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
	})
}
