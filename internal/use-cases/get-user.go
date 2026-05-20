package use_cases

import (
	"net/http"
	user "restapis-go/internal/entities"

	"restapis-go/internal/repository"

	"github.com/gin-gonic/gin"
)

// GetUsers godoc
// @Summary      Lista todos os usuários
// @Description  Retorna a lista de usuários
// @Tags         users
// @Produce      json
// @Success      200  {array}  user.User
// @Router       /users [get]
func GetUsers(c *gin.Context) {
	var users []user.User
	result := repository.Database.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}
