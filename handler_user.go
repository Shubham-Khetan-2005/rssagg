package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Shubham-Khetan-2005/rssagg/internal/auth"
	"github.com/Shubham-Khetan-2005/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiConfig *apiConfig)handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Name string `json:"name"`	
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	user, err := apiConfig.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID: uuid.New(),
		Name: params.Name,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(), 	
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create user: %v", err))
		return
	}	

	respondWithJSON(w, 201, databaseUserToUser(user))
	
}

func (apiConfig *apiConfig)handlerGetUser(w http.ResponseWriter, r *http.Request) {
	api_key, err := auth.GetAPIKey(r.Header)
	if err != nil {
		respondWithError(w, 403, fmt.Sprintf("Auth error: %v", err))
		return
	}

	user, err := apiConfig.DB.GetUserByAPIKey(r.Context(), api_key)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}