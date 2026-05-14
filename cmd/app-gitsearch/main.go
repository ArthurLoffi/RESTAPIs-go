package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type user struct {
	ID string `json:"ID"`
	Name string `json:"Nome"`
}

var users = []user{
	{ID: "1", Name: "Arthur"},
	{ID: "2", Name: "Caio"},
}

func getUsers(c *gin.Context)  {
	c.IndentedJSON(http.StatusOK, users)
}

func main() {
	router := gin.Default()
	router.GET("/users", getUsers)

	router.Run("localhost:8080")
}