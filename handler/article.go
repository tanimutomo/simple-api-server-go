package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tanimutomo/simple-api-server-go/db"
)

// Get a list of articles
func GetArticles() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		articles := db.GetArticles(username)
		c.JSON(http.StatusOK, gin.H{"articles": articles})
	}
}

// Post a new article
func PostArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")

		article := db.Article{Username: username}
		// Validation
		if err := c.Bind(&article); err != nil {
			BadRequestError(c, "Requested article is an invalid format")
		}

		db.PostArticle(article)
		c.JSON(http.StatusOK, article)
	}
}

// Update
func UpdateArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		articleIDStr := c.Param("articleID")

		// Check articleID compatibility
		articleID, err := strconv.Atoi(articleIDStr)
		if err != nil {
			NotFoundError(c, "articleID is invalid type. It should be uint.")
		}

		article := db.Article{Username: username}
		if err := c.Bind(&article); err != nil {
			BadRequestError(c, "Requested article is an invalid format")
		}
		if errmsg := db.UpdateArticle(articleID, article); errmsg != "" {
			BadRequestError(c, errmsg)
		}

		c.JSON(http.StatusOK, article)
	}
}

// Delete
func DeleteArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		articleIDStr := c.Param("articleID")

		// Check articleID compatibility
		articleID, err := strconv.Atoi(articleIDStr)
		if err != nil {
			NotFoundError(c, "articleID is invalid type. It should be uint.")
		}

		// Delete article
		if errmsg := db.DeleteArticle(articleID, username); errmsg != "" {
			BadRequestError(c, errmsg)
		}

		c.JSON(http.StatusOK, gin.H{"username": username, "articleID": articleID})
	}
}
