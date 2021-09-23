package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,longenough"`
}

var longenough validator.Func = func(fl validator.FieldLevel) bool {
	password, _ := fl.Field().Interface().(string)
	return len(password) > 4
}

var users = []User{}
var PORT = ":8008"

func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)
	c.String(http.StatusOK, "\n")
}

func addUser(c *gin.Context) {
	var newUser User

	err := c.ShouldBindJSON(&newUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.String(http.StatusBadRequest, "\n")
		return
	}

	users = append(users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
	c.String(http.StatusOK, "\n"+"User added!\n")
}

func main() {
	router := gin.Default()

	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		v.RegisterValidation("longenough", longenough)
	}

	router.GET("/get", getUsers)
	router.POST("/add", addUser)

	router.Run(PORT)
}
