package models

import (
	"time"
)

type Post struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}

var posts = []Post{
	{ID: 1, Title: "This is a post", Body: "Testing this first post", CreatedAt: time.Now()},
	{ID: 2, Title: "This is another post", Body: "Testing this second post", CreatedAt: time.Now()},
}

func GetAllPosts() ([]Post, error) {

	return posts, nil

}

func (p *Post) Save() error {
	var newPost = Post{p.ID, p.Title, p.Body, p.CreatedAt}
	posts = append(posts, newPost)
	// posts = append(posts, Post{len(posts) + 1}, "", "", time.Now())
	return nil
}

func (p *Post) Delete(id int) error {
	for index, post := range posts {
		if post.ID == id {
			posts = append(posts[:index], posts[index+1:]...)
		}
	}
	return nil
}
