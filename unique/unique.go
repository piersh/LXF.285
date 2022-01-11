package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

/*
Start Server in one terminal ($ symbolises user prompt):
$go run ./validate.go

Sample Usage in a second terminal ($ symbolises user prompt):
$curl -X GET localhost:8008/get
[]
$curl -X POST localhost:8008/add -H 'Content-Type: application/json' -d '{"username": "piersh", "password":  "Hello"}'
{
    "username": "piersh",
    "password": "Hello"
}
User added!
$curl -X POST localhost:8008/add -H 'Content-Type: application/json' -d '{"username": "piersh", "password":  "LXF12"}'
{"Error":"Key: 'User.Username' Error:Field validation for 'Username' failed on the 'unique' tag"}
$curl -X POST localhost:8008/add -H 'Content-Type: application/json' -d '{"username": "Linux", "password":  "Format"}'
{
    "username": "Linux",
    "password": "Format"
}
User added!
$curl -X GET localhost:8008/get                                                                 ✔ 
[
    {
        "username": "piersh",
        "password": "Hello"
    },
    {
        "username": "Linux",
        "password": "Format"
    }
]
$curl -X GET localhost:8008/get                                                                 ✔ 
[
    {
        "username": "piersh",
        "password": "Hello"
    },
    {
        "username": "Linux",
        "password": "Format"
    }
]
$curl -X POST localhost:8008/add -H 'Content-Type: application/json' -d '{"username": "piersh", "password":  "LXF"}'  
{"Error":"Key: 'User.Username' Error:Field validation for 'Username' failed on the 'unique' tag\nKey: 'User.Password' Error:Field validation for 'Password' failed on the 'longenough' tag"}
$curl -X POST localhost:8008/add -H 'Content-Type: application/json' -d '{"username": "piersh", "password":  ""}'   
{"Error":"Key: 'User.Username' Error:Field validation for 'Username' failed on the 'unique' tag\nKey: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag"}
$curl -X POST localhost:8008/add -H 'Content-Type: application/json' -d '{"username": "piersh"}'   
{"Error":"Key: 'User.Username' Error:Field validation for 'Username' failed on the 'unique' tag\nKey: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag"}
$curl -X POST localhost:8008/add -H 'Content-Type: application/json' -d '{"username": ""}'      ✔ 
{"Error":"Key: 'User.Username' Error:Field validation for 'Username' failed on the 'required' tag\nKey: 'User.Password' Error:Field validation for 'Password' failed on the 'required' tag"}
*/

type User struct {
	Username string `json:"username" binding:"required,unique"`
	Password string `json:"password" binding:"required,longenough"`
}

var longenough validator.Func = func(fl validator.FieldLevel) bool {
	password, _ := fl.Field().Interface().(string)
	return len(password) > 4
}

var users = []User{}
var PORT = ":8008"

var unique validator.Func = func(fl validator.FieldLevel) bool {
	username, _ := fl.Field().Interface().(string)
	for i := 0; i < len(users); i++ {
		if users[i].Username == username {
			return false
		}
	}
	return true
}


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
		v.RegisterValidation("unique", unique)
	}

	router.GET("/get", getUsers)
	router.POST("/add", addUser)

	router.Run(PORT)
}
