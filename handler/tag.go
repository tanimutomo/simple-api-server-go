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
		tags, errResp := db.GetTags(username)
		if errResp.IsError {
			SendErrorResponse(c, errResp.Status, errResp.Message)
		}
		c.JSON(http.StatusOK, gin.H{"tags": tags})
	}
}

// Add a new tag to existing article
func AddTag() gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.Param("username")
		articleIDStr := c.Param("articleID")

		// Check articleID compatibility
		articleID, err := strconv.ParseUint(articleIDStr, 10, 32)
		if err != nil {
			NotFoundError(c, "articleID is invalid type. It should be uint.")
		}
		tag := db.Tag{ArticleID: articleID}

		// Validation
		if err := c.Bind(&tag); err != nil {
			BadRequestError(c, "Requested tag is an invalid format")
		}

		// Insert a new tag to DB
		errResp := db.AddTag(tag, username)
		if errResp.IsError {
			SendErrorResponse(c, errResp.Status, errResp.Message)
		}

		c.JSON(http.StatusOK, tag)
	}
}
