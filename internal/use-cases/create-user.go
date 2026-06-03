package use_cases

import (
	"net/http"
	user "restapis-go/internal/entities"
	errorf "restapis-go/internal/error"
	db "restapis-go/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

// PostUser godoc
// @Summary      Adicionar um novo usuário
// @Description  Adiciona um novo usuário ao json
// @Tags         users
// @Produce      json
// @Security     BearerAuth
// @Success      201  {array}  user.User
// @Router       /post [post]
func CreateUser(name string, password string) (*user.User, *errorf.AppError){
	newUser := user.User{Name: name, Password: password}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		return &newUser, errorf.New(http.StatusInternalServerError, err.Error())
	}

	newUser.Password = string(hashedPassword)

	result := db.Database.Create(&newUser)
	if result.Error != nil {
		return nil, errorf.New(http.StatusInternalServerError, "User already exists")
	}

	return &newUser, nil
}
