package auth

import (
	"errors"
	"log"
	"net/http"
	"strings"
)

var ErrNoAuthHeaderIncluded = errors.New("no authorization header included")

func GetAPIKey(r *http.Request) (string, error) { // Изменение: аргумент *http.Request вместо http.Header
	log.Printf("Full headers received: %v", r.Header)
	authHeader := r.Header.Get("Authorization")
	log.Printf("Authorization header raw: %s", authHeader)
	if authHeader != "" {
		splitAuth := strings.Split(authHeader, " ")
		if len(splitAuth) < 2 || splitAuth[0] != "ApiKey" {
			log.Printf("Malformed authorization header: %s", authHeader)
			return "", errors.New("malformed authorization header")
		}
		log.Printf("Extracted API key from header: %s", splitAuth[1])
		return splitAuth[1], nil
	}

	// Fallback: читай из query ?api_key=<key>
	apiKey := r.URL.Query().Get("api_key")
	if apiKey == "" {
		log.Println("No authorization header or query param included")
		return "", ErrNoAuthHeaderIncluded
	}
	log.Printf("Extracted API key from query: %s", apiKey)
	return apiKey, nil
}
