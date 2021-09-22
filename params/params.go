package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var PORT = ":8008"

func timeHandler(c *gin.Context) {
	t := time.Now().Format(time.RFC1123)
	c.String(http.StatusOK, t)
}

func defaultHandler(c *gin.Context) {
	c.String(http.StatusNotFound, "Connecting from "+c.ClientIP())

}

func main() {
	fmt.Println("Starting HTTP server!")
	router := gin.Default()

	router.GET("/time", timeHandler)
	router.GET("/", timeHandler)
	router.NoRoute(defaultHandler)
	router.Run(PORT)
}
