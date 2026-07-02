package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()

	server.Run() // listens on 0.0.0.0:8080 by default
}
