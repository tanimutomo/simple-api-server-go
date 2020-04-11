package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tanimutomo/simple-api-server-go/auth"
	"github.com/tanimutomo/simple-api-server-go/db"
)

// Get a list of tags
func GetTags() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verrify Token
		_, err := auth.VerifyToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized"})
		} else {
			username := c.Param("username")
			tags := db.GetTags(username)
			c.JSON(http.StatusOK, gin.H{"tags": tags})
		}
	}
}

// Add a new tag to existing article
func AddTag() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Verrify Token
		_, err := auth.VerifyToken(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized"})
		} else {
			username := c.Param("username")
			articleIdStr := c.Param("articleId")
			articleId, err := strconv.ParseUint(articleIdStr, 10, 32)
			if err != nil {
				panic(err)
			}
			tag := db.Tag{ArticleID: articleId}
			// Validation
			if err := c.Bind(&tag); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"message": err})
			} else {
				if err := db.AddTag(tag, username); err != nil {
					c.JSON(http.StatusBadRequest, gin.H{"message": err})
				} else {
					c.JSON(http.StatusOK, gin.H{"message": "Success to add a new tag"})
				}
			}
		}
	}
}
