package utils

import (
	"crypto/rand"
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)


func HashJwt() (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "bar",
		"nbf": time.Date(2015, 10, 10, 12, 0, 0, 0, time.UTC).Unix(),
	})

	//* let's make another goroutine handle this simple task :p
	ch := make(chan []byte)
	go hashString(ch)
	hashSecret := <- ch

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(hashSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func hashString(ch chan []byte) {

	var str string = "nathapon:)1234"
	hashSecret := []byte(str)

	//* hash string to bytes
	_, err := rand.Read(hashSecret)
	if err != nil {
		log.Fatal(err)
	}

	ch <- hashSecret
}