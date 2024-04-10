package services

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var SECRET_KEY = os.Getenv("SECRET_KEY")

type SignedDetails struct {
	ID         primitive.ObjectID
	First_Name string
	Last_Name  string
	Email      string
	jwt.MapClaims
}

// TODO: change this from HS256 to RS256
func GenerateToken(id primitive.ObjectID, firstname string, lastname string, email string) (token string, err error) {
	claims := &SignedDetails{
		ID:         id,
		First_Name: firstname,
		Last_Name:  lastname,
		Email:      email,
		MapClaims: jwt.MapClaims{
			"iss": "Go Ecommerce",
			"exp": time.Now().Local().Add(time.Hour * time.Duration(24*7)).Unix(),
			"iat": time.Now().Local().Unix(),
		},
	}

	token, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		return "", err
	}

	return token, err
}

func ValidateToken(tokenString string) (claims *SignedDetails, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &SignedDetails{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(SECRET_KEY), nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*SignedDetails)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
