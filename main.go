package main

import (
	"github.com/gin-gonic/gin"
	"github.com/tanimutomo/simple-api-server-go/db"
	"github.com/tanimutomo/simple-api-server-go/handler"
)

func main() {
	r := gin.Default()

	db.Init()

	// Signup
	r.POST("/signup", handler.Signup())
	// Login
	r.POST("/login", handler.Login())

	// Get a list of tweets
	r.GET("/users/:username/articles", handler.GetArticles())
	// Post a new tweet
	r.POST("/users/:username/articles", handler.PostArticle())
	// Update
	r.POST("/users/:username/articles/:articleId", handler.UpdateArticle())
	// Delete
	r.DELETE("/users/:username/articles/:articleId", handler.DeleteArticle())

	r.Run(":8080")
}
