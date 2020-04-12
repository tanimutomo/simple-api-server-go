package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
	"github.com/tanimutomo/simple-api-server-go/auth"
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

	users := r.Group("/users", auth.VerifyToken())
	{
		// Get a list of articles
		users.GET("/:username/articles", handler.GetArticles())
		// Post a new article
		users.POST("/:username/articles", handler.PostArticle())
		// Update article
		users.POST("/:username/articles/:articleID", handler.UpdateArticle())
		// Add a new tag to article
		users.POST("/:username/articles/:articleID/tags", handler.AddTag())
		// Delete article
		users.DELETE("/:username/articles/:articleID", handler.DeleteArticle())

		// Get a list of tags
		users.GET("/:username/tags", handler.GetTags())
	}

	r.Run(":8080")
}
