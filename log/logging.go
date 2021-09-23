package main

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

var PORT = ":8008"
var LOGFILE = "/tmp/gin.log"

func timeHandler(c *gin.Context) {
	t := time.Now().Format(time.RFC1123)
	c.String(http.StatusOK, t)
}

func defaultHandler(c *gin.Context) {
	c.String(http.StatusNotFound, "Connecting from "+c.ClientIP())
}

func main() {
	gin.DisableConsoleColor()
	f, _ := os.Create(LOGFILE)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	router := gin.Default()
	router.GET("/time", timeHandler)
	router.PUT("/time", timeHandler)
	router.GET("/", timeHandler)
	router.NoRoute(defaultHandler)
	router.Run(PORT)
}
