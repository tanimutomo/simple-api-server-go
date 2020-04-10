package auth

import (
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	jwtrequest "github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
)

func GetToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := getToken()
		c.JSON(http.StatusOK, gin.H{"token": tokenString})
	}
}

func VerifyToken(request *http.Request) (*jwt.Token, error) {
	// Verrify signature
	return jwtrequest.ParseFromRequest(
		request, jwtrequest.OAuth2Extractor,
		func(token *jwt.Token) (interface{}, error) {
			b := []byte(os.Getenv("SIGNINGKEY"))
			return b, nil
		},
	)
}

func getToken() string {
	// set header
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims (json contents)
	claims := token.Claims.(jwt.MapClaims)
	claims["admin"] = true
	claims["name"] = "taro"
	claims["email"] = "taro@example.com"
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// signature
	tokenString, _ := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))

	return tokenString
}
