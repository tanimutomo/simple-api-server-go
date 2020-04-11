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
	// Login
	r.POST("/login", handler.Login())

	// Get a list of articles
	r.GET("/articles/:username", handler.GetArticles())
	// Post a new article
	r.POST("/articles/:username", handler.PostArticle())
	// Update
	r.POST("/articles/:username/:articleId", handler.UpdateArticle())
	// Delete
	r.DELETE("/articles/:username/:articleId", handler.DeleteArticle())

	r.Run(":8080")
}
