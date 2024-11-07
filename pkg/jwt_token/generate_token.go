package jwttoken

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
)

type Claims struct {
	UserId int    `json:"user_id"`
	UUID   string `json:"uuid"`
	jwt.StandardClaims
}

const SecretKey = "AFDAadwrwe&%&^DSDDSDI)(H:)!ss"

// GenerateJwtToken creates a JWT token with the userId and a generated UUID
func GenerateJwtToken(userId int) (string, error) {
	expirationTime := time.Now().Add(3 * time.Hour)
	newUUID := uuid.New().String()
	claims := &Claims{
		UserId: userId,
		UUID:   newUUID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		// Return the error properly
		return "", err
	}
	return tokenString, nil
}

// ValidateToken validates the JWT token and returns claims if valid
func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}
