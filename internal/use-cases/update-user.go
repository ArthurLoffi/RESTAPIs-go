package use_cases

import (
	"net/http"
	user "restapis-go/internal/entities"
	errorf "restapis-go/internal/error"
	db "restapis-go/internal/repository"
)

// SetUpdateUser godoc
// @Summary      Faz o update do user
// @Description  Muda o nome do usuário definido por url, com id e name
// @Tags         users
// @Produce      json
// @Security     BearerAuth
// @Success      200  {array}  user.User
// @Router       /update [patch]
func SetUpdateUser(name string, idString string) (user.User, *errorf.AppError) {
	var updatedUser user.User
	result := db.Database.Model(&user.User{}).Where("id = ?", idString).Updates(user.User{Name: name})
	
	// Se tiver erro vai retornar imediatamente
	if result.Error != nil {return updatedUser, errorf.New(http.StatusBadRequest, "Failed to update user")}

	result = db.Database.First(&updatedUser, idString)
	if result.Error != nil {return updatedUser, errorf.New(http.StatusNotFound, "User not found")}

	return updatedUser, nil
}