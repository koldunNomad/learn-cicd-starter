package auth

import (
	"errors"
	"net/http"
	"strings"
	"log"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

// GetAPIKey -
func GetAPIKey(headers http.Header) (string, error) {
	log.Printf("Headers received: %v", headers)
	authHeader := headers.Get("Authorization")
	if authHeader == "" {
		log.Println("No authorization header included")
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuth := strings.Split(authHeader, " ")
	if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
		log.Printf("Malformed authorization header: %s", authHeader)
		return "", errors.New("malformed authorization header")
	}
	log.Printf("Extracted API key: %s", splitAuth[1])
	return splitAuth[1], nil
}
