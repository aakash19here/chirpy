package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/aakash19here/chirpy/internal/database"
	"github.com/google/uuid"
)

type Chirp struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Body      string    `json:"body"`
	UserID    uuid.UUID `json:"user_id"`
}

func (cfg *apiConfig) getAllChirps(w http.ResponseWriter, r *http.Request) {
	c, err := cfg.dbQueries.GetAllChirps(r.Context())

	if err != nil {
		respondWithError(w, 500, "Error fetching chrips", err)
	}

	var chrips []Chirp

	for _, v := range c {
		chirp := Chirp{
			ID:        v.ID,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
			Body:      v.Body,
			UserID:    v.UserID,
		}

		chrips = append(chrips, chirp)
	}

	respondWithJSON(w, 200, chrips)

}

func (cfg *apiConfig) handlerChirps(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body   string    `json:"body"`
		UserId uuid.UUID `json:"user_id"`
	}

	type response struct {
		Chirp
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	cleanedOutput := censor(params.Body)

	dbChirp, err := cfg.dbQueries.CreateChirp(r.Context(), database.CreateChirpParams{
		ID:     uuid.New(),
		Body:   cleanedOutput,
		UserID: params.UserId,
	})

	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create chirp", err)
		return
	}

	respondWithJSON(w, http.StatusCreated, response{
		Chirp: Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			UserID:    dbChirp.UserID,
			Body:      dbChirp.Body,
		},
	})

	return

}

func censor(input string) string {
	// to lower -> "aakash is a dev"
	// split -> ["aakash", "is", "a", "dev"]
	// join -> "aakash is a dev"

	profaneWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}

	words := strings.Split(input, " ")

	for i, word := range words {
		replacement := strings.Repeat("*", 4)
		lowerCase := strings.ToLower(word)

		if _, ok := profaneWords[lowerCase]; ok {
			words[i] = replacement
		}
	}

	return strings.Join(words, " ")
}
