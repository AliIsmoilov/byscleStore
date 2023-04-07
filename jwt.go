package main

import (
	"fmt"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("secret")

type Credentials struct {
	Username string
	UserId   string
	Password string
}

type Claims struct {
	Username string
	UserId   string
	jwt.StandardClaims
}

func MakeJWT() string {

	var credentials Credentials
	credentials.Username = "superadmin"
	credentials.Password = "admin123"
	credentials.UserId = "1234567890"

	ExpiresAtTime := time.Now().Add(time.Hour * 20)

	claims := &Claims{
		Username: credentials.Username,
		UserId:   credentials.UserId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: ExpiresAtTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		panic(err)
	}

	return tokenString
}

func ParseJWT(tokenStr string) {

	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {

		if err == jwt.ErrSignatureInvalid {
			log.Println("Invalid token", err)
			return
		}
		panic(err)
	}

	fmt.Println(claims.Username)

	if !tkn.Valid {
		fmt.Println("Expired")
	}
}

func main() {

	fmt.Println(MakeJWT())
}
