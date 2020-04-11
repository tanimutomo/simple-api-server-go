package db

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/tanimutomo/simple-api-server-go/crypto"
)

func gormConnect() *gorm.DB {
	DBMS := os.Getenv("SASG_DBMS")
	USER := os.Getenv("SASG_USER")
	PASS := os.Getenv("SASG_PASS")
	DBNAME := os.Getenv("SASG_DBNAME")
	// postfix 'parse...' for charcode of mysql
	CONNECT := USER + ":" + PASS + "@/" + DBNAME + "?parseTime=true"

	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	return db
}

// Initialize DB
func Init() {
	db := gormConnect()

	defer db.Close()
	db.AutoMigrate(&Article{})
	db.AutoMigrate(&User{})
}

// Article //

// Declare article model
type Article struct {
	gorm.Model
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content" binding:"required"`
	Tag      string `json:"tag"`
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
		Tag:      article.Tag,
	})
}

// Update DB
func UpdateArticle(articleId int, updateArticle Article) interface{} {
	db := gormConnect()

	var article Article
	db.First(&article, articleId)
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

// User //

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required" gorm:"unique;not null"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

// Register a new user
func CreateUser(user User) interface{} {
	passwordEncrypt, _ := crypto.PasswordEncrypt(user.Password)
	db := gormConnect()
	defer db.Close()

	// Insert a new user to db
	if err := db.Create(
		&User{
			Username: user.Username,
			Password: passwordEncrypt,
			Email:    user.Email,
		},
	).Error; err != nil {
		return err
	}
	return nil
}

// Find a user
func GetUser(username string) (User, error) {
	db := gormConnect()
	var user User
	if err := db.First(&user, "username = ?", username).Error; err != nil {
		return user, err
	}
	db.Close()
	return user, nil
}

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
	if newArticle.Tag != "" {
		currentArticle.Tag = newArticle.Tag
	}
	return nil
}

func isUserMatched(usernameA string, usernameB string) bool {
	if usernameA != usernameB {
		return false
	}
	return true
}
