package routes

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/posts", getPosts)
	server.POST("/posts", createPost)
	server.DELETE("/posts/:id", deletePost)
}
