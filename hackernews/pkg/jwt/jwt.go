package jwt

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	SecretKey = []byte("secret")
)

/*
  GenerateToken
  the logics needs to be changed
  https://pkg.go.dev/github.com/golang-jwt/jwt#example-NewWithClaims-CustomClaimsType
*/
func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in generating key")
		return "", err
	}

	return tokenString, nil
}

/*
  ParseToken parses a jwt token and returns the username in it's claims
  https://pkg.go.dev/github.com/golang-jwt/jwt#example-Parse-Hmac
*/
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
