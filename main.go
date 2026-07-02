package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/events", getEvents)

	server.Run() // listens on 0.0.0.0:8080 by default
}

func getEvents(context *gin.Context) {
	context.JSON(http.StatusAccepted, gin.H{
		"message": "hello",
	})
}
