package utils

import (
	"addressbook/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// func to create TOKEN for user
func GenerateJWT(email string) (string, error) {
	secret := config.GetEnv("JWT_SECRET")
	//NewWithClaims creates a new [Token] with the specified signing method and claims. Additional options can be specified, but are currently unused
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{

		//Creates a new JWT token with claims (data inside the token):
		//"email": email — stores the user’s email inside the token.
		//signed using HMAC SHA-256, a secure signing algorithm.
		"email": email,
		"exp":   time.Now().Add(time.Hour * 1).Unix(),
	})
	//returns the string of the token which is signed
	return token.SignedString([]byte(secret))
}

// function for validation the signed token
func ValidateToken(tokenString string) (*jwt.Token, error) {
	secret := config.GetEnv("JWT_SECRET")
	return jwt.Parse(tokenString, func(token *jwt.Token) (any, error) {
		return []byte(secret), nil
	})
}
