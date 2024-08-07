package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/rileyarnie/rest-api/models"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/posts", models.GetAllPosts)
}
