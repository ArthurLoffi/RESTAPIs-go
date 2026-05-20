package use_cases

import (
	"net/http"
	user "restapis-go/internal/entities"

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
	c.JSON(http.StatusOK, user.Users)
}
