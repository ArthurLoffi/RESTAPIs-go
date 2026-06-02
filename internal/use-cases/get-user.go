package use_cases

import (
	"net/http"
	user "restapis-go/internal/entities"
	errorf "restapis-go/internal/error"
	"restapis-go/internal/repository"
)

// GetUsers godoc
// @Summary      Lista todos os usuários
// @Description  Retorna a lista de usuários
// @Tags         users
// @Produce      json
// @Success      200  {array}  user.User
// @Router       / [get]
func GetUsers() ([]user.User, *errorf.AppError) {
	var users []user.User
	result := repository.Database.Find(&users)
	if result.Error != nil {return users, errorf.New(http.StatusInternalServerError, "Can't get users")}

	return users, nil
}

// GetUserByID godoc
// @Summary      Lista o usuário com o ID desejado
// @Description  Retorna o usuário filtrado por ID
// @Tags         users
// @Produce      json
// @Success      200  {array}  user.User
// @Router       /:ID [get]
func GetUserByID(idString string) ([]user.User, *errorf.AppError) {
	var user []user.User
	result := repository.Database.First(&user, idString)
	if result.Error != nil {return user, errorf.New(http.StatusNotFound, "User not found")}

	return user, nil
}