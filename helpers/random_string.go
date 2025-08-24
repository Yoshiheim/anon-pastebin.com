package helpers

import (
	"math/rand"
	"time"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("okay_bro_@@"))

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano())) // Seed the random number generator

	result := make([]byte, length)
	for i := range result {
		result[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(result)
}
