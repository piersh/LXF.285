package main

import (
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
	path := c.FullPath()
	c.String(http.StatusNotFound, "From "+c.ClientIP()+" for "+path+"\n")

}

func getUsername(c *gin.Context) {
	username := c.Param("name")
	c.String(http.StatusOK, "Hello %s!\n", username)
}

func doSomething(c *gin.Context) {
	username := c.Param("name")
	doSomething := c.Param("doSomething")
	reply := username + " is visiting " + doSomething
	c.String(http.StatusOK, reply+"\n")
}

func main() {
	router := gin.Default()

	router.GET("/time", timeHandler)
	router.GET("/username/:name", getUsername)
	router.GET("/username/:name/*doSomething", doSomething)

	router.NoRoute(defaultHandler)
	router.Run(PORT)
}
