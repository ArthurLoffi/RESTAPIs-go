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
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil || body.Name == "" || body.Password == "" {
		errorf.Respond(c, errorf.New(http.StatusBadRequest, "invalid request body"))
		return
	}

	result, err := use_cases.AuthUser(use_cases.LoginInput{
		Name: body.Name,
		Password: body.Password,
	})
	
	if err != nil {
		errorf.Respond(c, err)
		return
	}

	claims := middleware.Claims{
		UserID: result.ID,
		User:   result.Name,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, signingErr := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if signingErr != nil {
		errorf.Respond(c, errorf.New(http.StatusInternalServerError, "failed to sign token"))
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Token": tokenString,
	})
}