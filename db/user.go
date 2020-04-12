package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/tanimutomo/simple-api-server-go/crypto"
)

type User struct {
	gorm.Model
	Username string `json:"username" binding:"required" gorm:"unique;not null"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type LoginUser struct {
	gorm.Model
	Username string `json:"username" binding:"required" gorm:"unique;not null"`
	Password string `json:"password" binding:"required"`
}

// Register a new user
func CreateUser(user User) string {
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
		return "Requested user is not compatible."
	}
	return ""
}

// Find a user
func GetUser(username string) (User, string) {
	db := gormConnect()
	var user User
	if err := db.First(&user, "username = ?", username).Error; err != nil {
		return user, "Requested user does not exists."
	}
	db.Close()
	return user, ""
}
