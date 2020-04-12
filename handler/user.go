package handler

import (
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
			BadRequestError(c, "Requested user is an invalid format")
		}

		// Check same username exists
		errResp := db.CreateUser(user)
		if errResp.IsError {
			SendErrorResponse(c, errResp.Status, errResp.Message)
		}

		c.JSON(http.StatusOK, user)
	}
}

// Login
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUser db.LoginUser
		// Validate request
		if err := c.Bind(&loginUser); err != nil {
			BadRequestError(c, "Requested login user is an invalid format")
		}

		// Check whether user is exists
		dbUser, errResp := db.GetUser(loginUser.Username)
		if errResp.IsError {
			SendErrorResponse(c, errResp.Status, errResp.Message)
		}

		// Compare sent password to db password
		dbPassword := dbUser.Password
		sentPassword := loginUser.Password
		if err := crypto.CompareHashAndPassword(
			dbPassword, sentPassword,
		); err != nil {
			UnauthorizedError(c, "Invalid password")
		}

		tokenString := GetToken(dbUser)
		c.JSON(http.StatusOK, gin.H{"user": dbUser, "token": tokenString})
	}
}
