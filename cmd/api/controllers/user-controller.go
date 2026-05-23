package controller

import (
	"net/http"
	use_cases "restapis-go/internal/use-cases"

	"github.com/gin-gonic/gin"
)

func NewUser(c *gin.Context) {
	name := c.Param("name")

	if name == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "User can't empty",
		})
		return
	}

	newUser, err := use_cases.CreateUser(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user": newUser,
	})
}

func ListUsers(c *gin.Context) {
	users, err := use_cases.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err,
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func ListUserByID(c *gin.Context) {
	idString := c.Param("ID")
	user, err := use_cases.GetUserByID(idString)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	idString := c.Param("ID")
	deletedUser, err := use_cases.DeleteUser(idString)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"Error": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
		"deleted": deletedUser,
	})
}