package main

import (
	"github.com/gin-gonic/gin"
	"github.com/rileyarnie/rest-api/routes"
)

func main() {
	server := gin.Default()
	routes.RegisterRoutes(server)
	server.Run()
}
