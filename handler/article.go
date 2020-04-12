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
		articles, errResp := db.GetArticles(username)
		if errResp.IsError {
			SendErrorResponse(c, errResp.Status, errResp.Message)
		}
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

		errResp := db.PostArticle(article)
		if errResp.IsError {
			SendErrorResponse(c, errResp.Status, errResp.Message)
		}
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

		// Update article contents
		errResp := db.UpdateArticle(articleID, article)
		if errResp.IsError {
			SendErrorResponse(c, errResp.Status, errResp.Message)
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
		errResp := db.DeleteArticle(articleID, username)
		if errResp.IsError {
			SendErrorResponse(c, errResp.Status, errResp.Message)
		}

		c.JSON(http.StatusOK, gin.H{"username": username, "articleID": articleID})
	}
}
