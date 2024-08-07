package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/", hello)
}

func hello(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"message": "Hello World"})

}
