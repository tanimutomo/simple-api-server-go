package auth

import (
	"net/http"
	"os"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	jwtrequest "github.com/dgrijalva/jwt-go/request"
	"github.com/tanimutomo/simple-api-server-go/db"
)

func VerifyToken(request *http.Request) (*jwt.Token, error) {
	// Verrify signature
	return jwtrequest.ParseFromRequest(
		request, jwtrequest.OAuth2Extractor,
		func(token *jwt.Token) (interface{}, error) {
			b := []byte(os.Getenv("SASG_SECRET"))
			return b, nil
		},
	)
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
