package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/tanimutomo/simple-api-server-go/db"
	"github.com/tanimutomo/simple-api-server-go/handler"
)

func main() {
	r := gin.Default()

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	db.Init()

	// Signup
	r.POST("/signup", handler.Signup())
	// Login -> Get token
	r.POST("/login", handler.Login())

	// Get a list of articles
	r.GET("/users/:username/articles", handler.GetArticles())
	// Post a new article
	r.POST("/users/:username/articles", handler.PostArticle())
	// Update article
	r.POST("/users/:username/articles/:articleID", handler.UpdateArticle())
	// Add a new tag to article
	r.POST("/users/:username/articles/:articleID/tags", handler.AddTag())
	// Delete article
	r.DELETE("/users/:username/articles/:articleID", handler.DeleteArticle())

	// Get a list of tags
	r.GET("/users/:username/tags", handler.GetTags())

	r.Run(":8080")
}
