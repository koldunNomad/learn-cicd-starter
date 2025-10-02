package main

import (
	"net/http"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
	"github.com/bootdotdev/learn-cicd-starter/internal/database"

	"log"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r)  // Изменение: передаём r
		if err != nil {
			log.Printf("Auth middleware: Failed to get API key: %v", err)
			respondWithError(w, http.StatusUnauthorized, "Couldn't find api key", err)
			return
		}
		log.Printf("Auth middleware: APIKey=%s", apiKey)
		user, err := cfg.DB.GetUser(r.Context(), apiKey)
		if err != nil {
			log.Printf("Auth middleware: Failed to get user for APIKey=%s: %v", apiKey, err)
			respondWithError(w, http.StatusNotFound, "Couldn't get user", err)
			return
		}
		log.Printf("Auth middleware: User found ID=%s, Name=%s, APIKey=%s", user.ID, user.Name, user.ApiKey)
		handler(w, r, user)
	}
}
