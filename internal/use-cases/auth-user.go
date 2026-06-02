package use_cases

import (
	"net/http"
	user "restapis-go/internal/entities"
	errorf "restapis-go/internal/error"
	"restapis-go/internal/repository"
)

// PostLogin godoc
// @Summary      Rota para logar na API
// @Description  Única rota pública para fazer login e receber o token
// @Tags         login
// @Produce      json
// @Security     BearerAuth
// @Success      201  {array}  user.User
// @Router       /login [post]
func AuthUser(name string) (*user.User, *errorf.AppError) {
	var user user.User
	result := repository.Database.Where("name = ?", name).First(&user)
	if result.Error != nil {
		return nil, errorf.New(http.StatusNotFound, "User don't exist")
	}
	return &user, nil
}