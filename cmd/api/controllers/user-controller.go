package controller

import (
	"net/http"

	errorf "restapis-go/internal/error"
	use_cases "restapis-go/internal/use-cases"
	"restapis-go/pkg"

	"github.com/gin-gonic/gin"
)

func NewUser(c *gin.Context) {
	var body map[string]string
	if err := c.ShouldBindJSON(&body); err != nil {
		errorf.Respond(c, errorf.New(http.StatusBadRequest, "invalid request body"))
		return
	}

	name := body["name"]
	if name == "" || !pkg.ValidateName(name) {
		errorf.Respond(c, errorf.New(http.StatusBadRequest, "invalid or missing name"))
		return
	}

	newUser, err := use_cases.CreateUser(name)
	if err != nil {
		errorf.Respond(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    newUser,
	})
}

func ListUsers(c *gin.Context) {
	users, err := use_cases.GetUsers()
	if err != nil {
		errorf.Respond(c, err)
		return
	}

	c.JSON(http.StatusOK, users)
}

func ListUserByID(c *gin.Context) {
	user, err := use_cases.GetUserByID(c.Param("ID"))
	if err != nil {
		errorf.Respond(c, err)
		return
	}

	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	deletedUser, err := use_cases.DeleteUser(c.Param("ID"))
	if err != nil {
		errorf.Respond(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
		"deleted": deletedUser,
	})
}

func UpdateUser(c *gin.Context) {
	var body map[string]string
	if err := c.ShouldBindJSON(&body); err != nil {
		errorf.Respond(c, errorf.New(http.StatusBadRequest, "invalid request body"))
		return
	}

	name := body["name"]
	if name == "" || !pkg.ValidateName(name) {
		errorf.Respond(c, errorf.New(http.StatusBadRequest, "invalid or missing name"))
		return
	}

	updatedUser, err := use_cases.SetUpdateUser(name, c.Param("ID"))
	if err != nil {
		errorf.Respond(c, err)
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}