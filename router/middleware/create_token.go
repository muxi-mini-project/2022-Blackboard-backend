package middleware

import (
	"blackboard/pkg/token"
	"github.com/dgrijalva/jwt-go"
	"time"
)

func CreateToken(ID string) (string, error) {
	newWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, &token.Jwt{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(200 * time.Hour).Unix(),
			Issuer:    "Wishforpeace",
			IssuedAt:  time.Now().Unix(),
		},
		StudentID: ID,
	})
	var Secret = []byte("blackboard")
	return newWithClaims.SignedString(Secret)
}
