package use_cases

import (
	"net/http"
	user "restapis-go/internal/entities"
	errorf "restapis-go/internal/error"
	db "restapis-go/internal/repository"
)

// DeleteUser godoc
// @Summary      Remover usuário por id
// @Description  Remove o usuário com o determinado id no json
// @Tags         users
// @Produce      json
// @Success      200  {array}  user.User
// @Router       /delete/:ID [delete]
func DeleteUser(idString string) (user.User, *errorf.AppError){
	var user user.User
	result := db.Database.First(&user, idString)

	// Se tiver erro retorna antes de passar para a outra query da DB
	if result.Error != nil {return user, errorf.New(http.StatusNotFound, "User not found")}

	result = db.Database.Delete(&user)
	if result.Error != nil {return user, errorf.New(http.StatusInternalServerError, "It was not possible to delete the user")}
	return user, nil
}
