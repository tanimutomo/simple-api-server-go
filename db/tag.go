package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Tag struct {
	gorm.Model
	Name      string
	ArticleID uint64
}

// Get all tags
func GetTags(username string) []Tag {
	db := gormConnect()

	defer db.Close()
	var tags []Tag
	db.Table("tags").Select("tags.name").Joins("left join articles on tags.article_id = articles.id").Find(&tags)
	return tags
}

// Add tags to article
func AddTag(tag Tag, username string) interface{} {
	db := gormConnect()

	// Check the requested tag is already exists
	defer db.Close()
	var count int
	db.Table("tags").Joins("left join articles on tags.article_id = articles.id").Where("articles.username = ? AND articles.id = ? AND tags.name = ?", username, tag.ArticleID, tag.Name).Count(&count)
	if count != 0 {
		return "This article has already the requested tag. Can't give same tags to one article."
	}

	// Add a new tag
	db.Create(&Tag{
		Name:      tag.Name,
		ArticleID: tag.ArticleID,
	})
	return nil
}
