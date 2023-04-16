package helpers

import (
	"fmt"
	model "chapter3-sesi2/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type MyClaims struct {
	User *model.User
	jwt.StandardClaims
}

const secretKey = "rahasia"

func HashPassword(plaintPassword string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plaintPassword), 14)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func PasswordIsMatch(plaintPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plaintPassword))
	if err != nil {
		return false
	}

	return true
}

func GenerateToken(myClaims MyClaims) (string, error) {
	tokenExpirationTime := time.Now().Add(time.Hour * 1)
	claims := MyClaims{
		User: myClaims.User,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: tokenExpirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(secretKey))

	return signedToken, err
}

func VerifyAccessToken(tokenString string) (*MyClaims, error) {
	claims := &MyClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Token is invalid")
	}

	claims, ok := token.Claims.(*MyClaims)

	if !ok {
		return nil, fmt.Errorf("Couldn't parse claims")
	}

	return claims, nil
}
