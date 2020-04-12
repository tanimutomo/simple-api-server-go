package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanimutomo/simple-api-server-go/auth"
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
			c.Abort()
		}

		// Check same username exists
		if err := db.CreateUser(user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
			c.Abort()
		}

		c.JSON(http.StatusFound, gin.H{"message": "Success to signup"})
	}
}

// Login
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUser db.LoginUser
		// Validate request
		if err := c.Bind(&loginUser); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
			c.Abort()
		}

		// Check whether user is exists
		dbUser, err := db.GetUser(loginUser.Username)
		if err != nil {
			log.Println("Failed to login")
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
			c.Abort()
		}

		// Compare sent password to db password
		dbPassword := dbUser.Password
		sentPassword := loginUser.Password
		if err := crypto.CompareHashAndPassword(
			dbPassword, sentPassword,
		); err != nil {
			log.Println("Failed to login")
			c.JSON(http.StatusBadRequest, gin.H{"Error": err})
			c.Abort()
		}

		log.Println("Success to login")
		tokenString := auth.GetToken(dbUser)
		c.JSON(
			http.StatusFound,
			gin.H{"message": "Success to login", "token": tokenString},
		)
	}
}
