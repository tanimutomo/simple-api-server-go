package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tanimutomo/simple-api-server-go/crypto"
	"github.com/tanimutomo/simple-api-server-go/db"
)

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Signup
func Signup() gin.HandlerFunc {
	return func(c *gin.Context) {
		var user db.User
		// Validation
		if err := c.Bind(&user); err != nil {
			BadRequestError(c, "Requested user is an invalid format")
		}

		// Check same username exists
		if err := db.CreateUser(user); err != nil {
			switch e := err.(type) {
			case *db.ErrorResponse:
				SendErrorResponse(c, e.Status, e.Message)
			default:
				InternalServerError(c, "Unknown Type Error")
			}
		}

		c.JSON(http.StatusOK, user)
	}
}

// Login
func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginUser LoginRequest

		// Validate request
		if err := c.Bind(&loginUser); err != nil {
			BadRequestError(c, "Requested login user is an invalid format")
		}

		// Check whether user is exists
		dbUser, err := db.GetUser(loginUser.Username)
		if err != nil {
			switch e := err.(type) {
			case *db.ErrorResponse:
				SendErrorResponse(c, e.Status, e.Message)
			default:
				InternalServerError(c, "Unknown Type Error")
			}
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
		c.JSON(http.StatusOK, gin.H{
			"username": loginUser.Username,
			"token":    tokenString,
		})
	}
}
