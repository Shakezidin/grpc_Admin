package JWT

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var jwtKey = []byte("SECRETKEY")

type jwtClaim struct {
	Email string
	Role  string
	jwt.StandardClaims
}

func GenerateJWT(email,role string) (string, error) {
	expirationTime := time.Now().Add(1 * time.Hour)
	claims := &jwtClaim{
		Email: email,
		Role: role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
