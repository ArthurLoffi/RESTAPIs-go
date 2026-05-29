package controller

import (
	"net/http"
	"os"
	errorf "restapis-go/internal/error"
	"restapis-go/internal/middleware"
	use_cases "restapis-go/internal/use-cases"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(c *gin.Context) {
	var body struct {
		Name string `json:"name"`
	}

	if err := c.ShouldBindJSON(&body);err != nil || body.Name == ""{
		e := http.StatusBadRequest
		c.JSON(e, errorf.FormatedError(e, "EOF"))
		return
	}

	result, err := use_cases.AuthUser(body.Name)
	if err != nil {
		e := http.StatusUnauthorized
		c.JSON(e, errorf.FormatedError(e, err.Error()))
		return
	}

	claims := middleware.Claims{
		UserID: result.ID,
		User: result.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		e := http.StatusInternalServerError
		c.JSON(e, errorf.FormatedError(e, err.Error()))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Token": tokenString,
	})
}