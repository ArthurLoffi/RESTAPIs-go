package use_cases

import (
	"net/http"
	db "restapis-go/internal/repository"
	user "restapis-go/internal/entities"

	"github.com/gin-gonic/gin"
)

// DeleteUser godoc
// @Summary      Remover usuário por id
// @Description  Remove o usuário com o determinado id no json
// @Tags         users
// @Produce      json
// @Success      200  {array}  user.User
// @Router       /users/delete/:ID [delete]
func DeleteUser(c *gin.Context) {
	idString := c.Param("ID")

	var user user.User
	result := db.Database.First(&user, idString)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "user not found",
		})
		return
	}

	db.Database.Delete(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully!",
	})
}
