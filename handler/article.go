package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tanimutomo/simple-api-server-go/auth"
	"github.com/tanimutomo/simple-api-server-go/db"
)

// Get a list of articles
func GetArticles() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verrify Token
		_, err := auth.VerifyToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized"})
		} else {
			username := c.Param("username")
			articles := db.GetArticles(username)
			c.JSON(http.StatusOK, gin.H{"articles": articles})
		}
	}
}

// Post a new article
func PostArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verrify Token
		_, err := auth.VerifyToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized"})
		} else {
			username := c.Param("username")
			article := db.Article{Username: username}
			// Validation
			if err := c.Bind(&article); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err, "article": article})
			} else {
				db.PostArticle(article)
				c.JSON(http.StatusOK, gin.H{"message": "Success to post a new article"})
			}
		}
	}
}

// Update
func UpdateArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verrify Token
		_, err := auth.VerifyToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized"})
		} else {
			username := c.Param("username")
			articleId := c.Param("articleId")
			aid, err := strconv.Atoi(articleId)
			if err != nil {
				panic(err)
			}
			article := db.Article{Username: username}
			c.Bind(&article)
			if err := db.UpdateArticle(aid, article); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"Error": err})
			} else {
				c.JSON(http.StatusOK, gin.H{"message": "Sccess to update a article"})
			}
		}
	}
}

// Delete
func DeleteArticle() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verrify Token
		_, err := auth.VerifyToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized"})
		} else {
			username := c.Param("username")
			articleId := c.Param("articleId")
			aid, err := strconv.Atoi(articleId)
			if err != nil {
				panic(err)
			}
			if err := db.DeleteArticle(aid, username); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"Error": err})
			} else {
				c.JSON(http.StatusFound, gin.H{"message": "Success to delete a article"})
			}
		}
	}
}
