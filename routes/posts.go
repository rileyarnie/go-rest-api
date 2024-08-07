package routes

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rileyarnie/rest-api/models"
)

func getPosts(context *gin.Context) {
	posts, err := models.GetAllPosts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"posts": posts})

}

func createPost(context *gin.Context) {
	var newPost models.Post
	posts, _ := models.GetAllPosts()
	err := context.ShouldBindJSON(&newPost)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Couldn't parse data"})
	}
	newPost.ID = len(posts) + 1
	newPost.CreatedAt = time.Now()
	newPost.Save()
}

func deletePost(context *gin.Context) {
	var post models.Post
	var postId, err = strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't parse post id"})
	}
	err = post.Delete(int(postId))
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "couldn't delete post id"})
	}

	context.JSON(http.StatusOK, gin.H{"message": "successfully delete post"})
}
