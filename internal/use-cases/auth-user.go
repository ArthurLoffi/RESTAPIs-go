package use_cases

import (
	"net/http"
	user "restapis-go/internal/entities"
	errorf "restapis-go/internal/error"
	"restapis-go/internal/repository"

	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Name string
	Password string
}

// PostLogin godoc
// @Summary      Rota para logar na API
// @Description  Única rota pública para fazer login e receber o token
// @Tags         login
// @Produce      json
// @Security     BearerAuth
// @Success      201  {array}  user.User
// @Router       /login [post]
func AuthUser(login LoginInput) (*user.User, *errorf.AppError) {
	var user user.User

	result := repository.Database.Where("name = ?", login.Name).First(&user)
	if result.Error != nil {
		return nil, errorf.New(http.StatusUnauthorized, "Invalid credentials")
	}

	// Passa para hash a senha
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password)); err != nil {
		return nil, errorf.New(http.StatusUnauthorized, "Invalid credentials")
	}

	return &user, nil
}