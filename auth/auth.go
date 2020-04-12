package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	jwtrequest "github.com/dgrijalva/jwt-go/request"
	"github.com/tanimutomo/simple-api-server-go/db"
)

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		_, err := jwtrequest.ParseFromRequest(
			c.Request, jwtrequest.OAuth2Extractor,
			func(token *jwt.Token) (interface{}, error) {
				b := []byte(os.Getenv("SASG_SECRET"))
				return b, nil
			},
		)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Unauthorized"})
			c.Abort()
		}
	}
}

func GetToken(user db.User) string {
	// set header
	token := jwt.New(jwt.SigningMethodHS256)

	// set claims (json contents)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["email"] = user.Email
	claims["iat"] = time.Now()
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	// signature
	tokenString, _ := token.SignedString([]byte(os.Getenv("SASG_SECRET")))

	return tokenString
}
