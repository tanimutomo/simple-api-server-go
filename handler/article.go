package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanimutomo/simple-api-server-go/article"
	"github.com/tanimutomo/simple-api-server-go/auth"
)

func ArticleGet(articles *article.Articles) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verrify Token
		_, err := auth.VerifyToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err})
		}

		result := articles.GetAll()
		c.JSON(http.StatusOK, result) // Return all articles
	}
}

func ArticlePost(articles *article.Articles) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verrify Token
		_, err := auth.VerifyToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"message": err})
		}

		requestBody := article.Item{}
		c.Bind(&requestBody) // Push requested article to requestBody by c.Bind

		item := article.Item{
			Title:       requestBody.Title,
			Description: requestBody.Description,
		}
		articles.Add(item) // Add new article to articles in main.go

		c.JSON(http.StatusOK, gin.H{"message": "Success to post a new article"}) // Return ack
	}
}
