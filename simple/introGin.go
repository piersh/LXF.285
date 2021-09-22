package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

var PORT = ":8008"

func main() {
	fmt.Println("Starting HTTP server!")
	router := gin.Default()

	router.Run(PORT)
}
