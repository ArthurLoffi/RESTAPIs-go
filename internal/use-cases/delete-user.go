package use_cases

import (
	db "restapis-go/internal/repository"
	user "restapis-go/internal/entities"
)

// DeleteUser godoc
// @Summary      Remover usuário por id
// @Description  Remove o usuário com o determinado id no json
// @Tags         users
// @Produce      json
// @Success      200  {array}  user.User
// @Router       /delete/:ID [delete]
func DeleteUser(idString string) (user.User, error){
	var user user.User
	result := db.Database.First(&user, idString)

	// Se tiver erro retorna antes de passar para a outra query da DB
	if result.Error != nil {return user, result.Error}

	result = db.Database.Delete(&user)
	return user, result.Error
}
