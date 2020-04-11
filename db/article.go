package db

import (
	// "log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Tags     []Tag  `json:"tags" gorm:"foreignkey:ArticleID"`
	Username string `json:"username"`
}

// Get all data
func GetArticles(username string) []Article {
	db := gormConnect()

	defer db.Close()
	var articles []Article
	// Get all article data by specifying empty condition as the Find argument
	db.Order("created_at desc").Where("Username = ?", username).Find(&articles)
	return articles
}

// Insert data
func PostArticle(article Article) {
	db := gormConnect()

	defer db.Close()
	db.Create(&Article{
		Title:    article.Title,
		Content:  article.Content,
		Username: article.Username,
		Tags:     article.Tags,
	})
}

// Update DB
func UpdateArticle(articleID int, updateArticle Article) interface{} {
	db := gormConnect()

	// Delete old tags associated to this article
	var tag Tag
	db.Where("article_id = ?", articleID).Delete(&tag)

	// Update article
	var article Article
	db.First(&article, articleID)
	if err := updateArticleContents(&article, updateArticle); err != nil {
		return err
	}
	db.Save(&article)

	db.Close()
	return nil
}

// Delete a article
func DeleteArticle(id int, username string) interface{} {
	db := gormConnect()
	var article Article
	if err := db.First(&article, id).Error; err != nil {
		return "The article is not existed."
	}
	if !isUserMatched(article.Username, username) {
		return "User is not matched"
	}
	db.Delete(&article)
	db.Close()
	return nil
}

// Utility functions //

func updateArticleContents(currentArticle *Article, newArticle Article) interface{} {
	if !isUserMatched(currentArticle.Username, newArticle.Username) {
		return "User is not matched"
	}
	if newArticle.Title != "" {
		currentArticle.Title = newArticle.Title
	}
	if newArticle.Content != "" {
		currentArticle.Content = newArticle.Content
	}
	if len(newArticle.Tags) != 0 {
		currentArticle.Tags = newArticle.Tags
	}
	return nil
}

func isUserMatched(usernameA string, usernameB string) bool {
	if usernameA != usernameB {
		return false
	}
	return true
}
