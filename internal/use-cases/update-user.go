package use_cases

import (
	user "restapis-go/internal/entities"
	db "restapis-go/internal/repository"
)

// SetUpdateUser godoc
// @Summary      Faz o update do user
// @Description  Muda o nome do usuário definido por url, com id e name
// @Tags         users
// @Produce      json
// @Success      200  {array}  user.User
// @Router       /update [patch]
func SetUpdateUser(name string, idString string) (user.User, error) {
	var updatedUser user.User
	result := db.Database.Model(&user.User{}).Where("id = ?", idString).Updates(user.User{Name: name})
	
	// Se tiver erro vai retornar imediatamente
	if result.Error != nil {return updatedUser, result.Error}

	result = db.Database.First(&updatedUser, idString)
	return updatedUser, result.Error
}