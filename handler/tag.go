package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tanimutomo/simple-api-server-go/db"
)

// Get a list of tags
func GetTags() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		tags := db.GetTags(username)
		c.JSON(http.StatusOK, gin.H{"tags": tags})
	}
}

// Add a new tag to existing article
func AddTag() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		articleIDStr := c.Param("articleID")

		articleID, err := strconv.ParseUint(articleIDStr, 10, 32)
		if err != nil {
			panic(err)
		}
		tag := db.Tag{ArticleID: articleID}

		// Validation
		if err := c.Bind(&tag); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
			c.Abort()
		}

		// Insert a new tag to DB
		if err := db.AddTag(tag, username); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err})
			c.Abort()
		}

		c.JSON(http.StatusOK, gin.H{"message": "Success to add a new tag"})
	}
}
