package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var PORT = ":8008"

func main() {
	fmt.Println("Starting HTTP server!")
	router := gin.Default()

	router.GET("/time", time)
	router.POST("/", default)

	router.Run(PORT)
}
