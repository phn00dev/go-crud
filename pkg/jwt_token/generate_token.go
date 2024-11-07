package jwttoken

import (
	"time"

	"github.com/dgrijalva/jwt-go"

)

type Claims struct {
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

const SecretKey = "AFDAadwrwe&%&^DSDDSDI)(H:)!ss"

func GenerateJwtToken(userId int, username string) (string, error) {
	expirationTime := time.Now().Add(3 * time.Hour)
	claims := &Claims{
		UserId:   userId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", nil
	}
	return tokenString, nil
}
