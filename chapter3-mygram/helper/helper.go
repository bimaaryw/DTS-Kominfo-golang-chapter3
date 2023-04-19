package helper

import (
	"fmt"
	"os"
	"time"

	"chapter3-mygram/model"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

var validate *validator.Validate

type authCustomClaims struct {
	Username string
	UserID   string
	jwt.StandardClaims
}

func GenerateID() string {
	return uuid.New().String()
}

func Hash(plain string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(plain), 8)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func PasswordIsMatch(plaintPassword, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plaintPassword))
	if err != nil {
		return false
	}

	return true
}

func GenerateToken(userData model.User) (string, error) {
	claims := &authCustomClaims{
		Username: userData.Username,
		UserID:   userData.UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 1).Unix(),
			Issuer:    "myGram",
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	tokeStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokeStr, nil
}

func VerifyAccessToken(tokenStr string) (*authCustomClaims, error) {
	claims := &authCustomClaims{}
	token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("Token is invalid")
	}

	claims, ok := token.Claims.(*authCustomClaims)

	if !ok {
		return nil, fmt.Errorf("Couldn't parse claims")
	}

	return claims, nil
}
