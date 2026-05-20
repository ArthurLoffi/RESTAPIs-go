package use_cases

import (
	"net/http"
	user "restapis-go/internal/entities"
	"strconv"

	"github.com/gin-gonic/gin"
)

// DeleteUser godoc
// @Summary      Remover usuário por id
// @Description  Remove o usuário com o determinado id no json
// @Tags         users
// @Produce      json
// @Success      200  {array}  user.User
// @Router       /users/delete [delete]
func DeleteUser(c *gin.Context) {
	idString := c.Param("ID")

	id, err := strconv.ParseUint(idString, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid ID",
		})
	}

	for i, u := range user.Users {
		if u.ID == uint(id) {
			user.Users = append(user.Users[:i], user.Users[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "Deleted user sucessfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}
