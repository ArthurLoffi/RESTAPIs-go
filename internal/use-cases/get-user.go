package use_cases

import (
	user "restapis-go/internal/entities"
	"restapis-go/internal/repository"
)

// GetUsers godoc
// @Summary      Lista todos os usuários
// @Description  Retorna a lista de usuários
// @Tags         users
// @Produce      json
// @Success      200  {array}  user.User
// @Router       / [get]
func GetUsers() ([]user.User, error) {
	var users []user.User
	result := repository.Database.Find(&users)

	return users, result.Error
}

// GetUserByID godoc
// @Summary      Lista o usuário com o ID desejado
// @Description  Retorna o usuário filtrado por ID
// @Tags         users
// @Produce      json
// @Success      200  {array}  user.User
// @Router       /:ID [get]
func GetUserByID(idString string) ([]user.User, error) {
	var users []user.User
	result := repository.Database.First(&users, idString)

	return users, result.Error
}