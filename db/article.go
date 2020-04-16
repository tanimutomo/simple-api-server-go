package db

import (
	"net/http"

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
func GetArticles(username string) ([]Article, error) {
	var articles []Article

	db, err := gormConnect()
	if err != nil {
		return articles, err
	}

	defer db.Close()
	// Get all article data by specifying empty condition as the Find argument
	db.Order("created_at desc").Where("Username = ?", username).Find(&articles)
	return articles, nil
}

// Insert data
func PostArticle(article Article) error {
	db, err := gormConnect()
	if err != nil {
		return err
	}

	defer db.Close()
	db.Create(&Article{
		Title:    article.Title,
		Content:  article.Content,
		Username: article.Username,
		Tags:     article.Tags,
	})

	return nil
}

// Update DB
func UpdateArticle(articleID int, updateArticle Article) error {
	db, err := gormConnect()
	if err != nil {
		return err
	}

	// Delete old tags associated to this article
	var tag Tag
	db.Where("article_id = ?", articleID).Delete(&tag)

	// Check user is compatible
	var article Article
	db.First(&article, articleID)
	if !isUserMatched(article.Username, updateArticle.Username) {
		return &ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "User is not matched",
		}
	}

	updateArticleContents(&article, updateArticle)
	db.Save(&article)
	db.Close()

	return nil
}

// Delete a article
func DeleteArticle(id int, username string) error {
	db, err := gormConnect()
	if err != nil {
		return err
	}

	var article Article
	if err := db.First(&article, id).Error; err != nil {
		return &ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "The article is not existed.",
		}
	}
	if !isUserMatched(article.Username, username) {
		return &ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "User is not matched",
		}
	}

	db.Delete(&article)
	db.Close()

	return nil
}

// Utility functions //

func updateArticleContents(currentArticle *Article, newArticle Article) interface{} {
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
	return usernameA != usernameB
}
