package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

var users = []User{}
var PORT = ":8008"

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
	c.String(http.StatusOK, "\n")
}

func addUser(c *gin.Context) {
	var newUser User

	// BindJSON binds the received JSON to User structure
	err := c.BindJSON(&newUser)
	if err != nil {
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
	c.String(http.StatusOK, "\n"+"User added!\n")
}

func main() {
	router := gin.Default()
	router.GET("/get", getUsers)
	router.POST("/add", addUser)

	router.Run(PORT)
}
