package handler

import (
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	jwtrequest "github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
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
			UnauthorizedError(c, "Invalid token. "+err.Error())
			return
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
