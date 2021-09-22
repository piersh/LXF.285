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

func getUsername(c *gin.Context) {
	username := c.Param("name")
	fmt.Println("Username:", username)
	c.String(http.StatusOK, "Hello %s!", username)
}

func main() {
	fmt.Println("Starting HTTP server!")
	router := gin.Default()

	router.GET("/time", timeHandler)
	router.NoRoute(defaultHandler)

	router.GET("/username/:name", getUsername)

	router.Run(PORT)
}
