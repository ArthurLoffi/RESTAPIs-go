package use_cases

import (
	user "restapis-go/internal/entities"
	db "restapis-go/internal/repository"
)

// PostUser godoc
// @Summary      Adicionar um novo usuário
// @Description  Adiciona um novo usuário ao json
// @Tags         users
// @Produce      json
// @Success      201  {array}  user.User
// @Router       /users/post/:name [post]
func CreateUser(name string) (*user.User, error){
	newUser := user.User{Name: name}

	result := db.Database.Create(&newUser)
	if result.Error != nil {
		return nil, result.Error
	}

	return &newUser, nil
}
