package use_cases

import (
	"net/http"
	user "restapis-go/internal/entities"
	errorf "restapis-go/internal/error"
	db "restapis-go/internal/repository"
)

// PostUser godoc
// @Summary      Adicionar um novo usuário
// @Description  Adiciona um novo usuário ao json
// @Tags         users
// @Produce      json
// @Success      201  {array}  user.User
// @Router       /post [post]
func CreateUser(name string) (*user.User, *errorf.AppError){
	newUser := user.User{Name: name}

	result := db.Database.Create(&newUser)
	if result.Error != nil {
		return nil, errorf.New(http.StatusInternalServerError, "User already exists")
	}

	return &newUser, nil
}
