package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanimutomo/simple-api-server-go/crypto"
	"github.com/tanimutomo/simple-api-server-go/db"
)

// Signup
func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user db.User
		// Validation
		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		} else {
			// Check same username exists
			if err := db.CreateUser(user); err != nil {
				c.JSON(http.StatusBadRequest, gin.H{"Error": err})
			} else {
				c.JSON(http.StatusFound, gin.H{"message": "Success to signup"})
			}
		}
	}
}

// Login
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user db.User
		if err := c.Bind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
		} else {
			// Get hashed password
			if existingUser, err := db.GetUser(user.Username); err != nil {
				log.Println("Failed to login")
				c.JSON(http.StatusBadRequest, gin.H{"Error": err})
			} else {
				dbPassword := existingUser.Password
				sentPassword := user.Password
				// Compare user sent password to db password
				if err := crypto.CompareHashAndPassword(
					dbPassword, sentPassword,
				); err != nil {
					log.Println("Failed to login")
					c.JSON(http.StatusBadRequest, gin.H{"Error": err})
				} else {
					log.Println("Success to login")
					c.JSON(http.StatusFound, gin.H{"message": "Success to login"})
				}
			}
		}
	}
}
