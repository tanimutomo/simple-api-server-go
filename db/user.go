package db

import (
	"net/http"

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

// Register a new user
func CreateUser(user User) error {
	passwordEncrypt, _ := crypto.PasswordEncrypt(user.Password)
	db, err := gormConnect()
	if err != nil {
		return err
	}

	defer db.Close()

	// Insert a new user to db
	if err := db.Create(
		&User{
			Username: user.Username,
			Password: passwordEncrypt,
			Email:    user.Email,
		},
	).Error; err != nil {
		return &ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Requested user is not compatible.",
		}
	}
	return nil
}

// Find a user
func GetUser(username string) (User, error) {
	var user User

	db, err := gormConnect()
	if err != nil {
		return user, err
	}

	if err := db.First(&user, "username = ?", username).Error; err != nil {
		return user, &ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "Requested user does not exists.",
		}
	}
	db.Close()
	return user, nil
}
