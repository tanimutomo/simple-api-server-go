package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"net/http"
)

type Tag struct {
	gorm.Model
	Name      string
	ArticleID uint64
}

// Get all tags
func GetTags(username string) ([]Tag, ErrorResponse) {
	var tags []Tag

	db, errResp := gormConnect()
	if errResp.IsError {
		return tags, errResp
	}

	defer db.Close()
	db.Table("tags").Select("tags.name").Joins("left join articles on tags.article_id = articles.id").Find(&tags)
	return tags, ErrorResponse{IsError: false}
}

// Add tags to article
func AddTag(tag Tag, username string) ErrorResponse {
	db, errResp := gormConnect()
	if errResp.IsError {
		return errResp
	}

	// Check the requested tag is already exists
	defer db.Close()
	var count int
	db.Table("tags").Joins("left join articles on tags.article_id = articles.id").Where("articles.username = ? AND articles.id = ? AND tags.name = ?", username, tag.ArticleID, tag.Name).Count(&count)
	if count != 0 {
		return ErrorResponse{
			IsError: true,
			Status:  http.StatusBadRequest,
			Message: "This article has already the requested tag. Can't give same tags to one article.",
		}
	}

	// Add a new tag
	db.Create(&Tag{
		Name:      tag.Name,
		ArticleID: tag.ArticleID,
	})
	return ErrorResponse{IsError: false}
}
