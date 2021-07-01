package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	app := gin.Default()
	gin.SetMode(gin.DebugMode) //change to gin.ReleaseMode

	app.GET("/", func(c *gin.Context) {

		c.JSON(http.StatusOK, gin.H{"msg": "Hello from gin!", "type": c.ContentType()})
		c.Status(200)

	})

	app.Run(":5000")
}
