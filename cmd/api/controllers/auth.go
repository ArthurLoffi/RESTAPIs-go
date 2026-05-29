package controller

import (
	"net/http"
	"os"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Name should not be null",
		})
		return
	}

	result, err := use_cases.AuthUser(body.Name)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": err.Error(),
		})
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
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Token": tokenString,
	})
}