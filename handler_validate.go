package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerChirpsValidate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}
	type returnVals struct {
		CleanBody string `json:"cleaned_body"`
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

	respondWithJSON(w, http.StatusOK, returnVals{
		CleanBody: cleanedOutput,
	})
}

func censor(input string) string {
	// to lower -> "aakash is a dev"
	// split -> ["aakash", "is", "a", "dev"]
	// join -> "aakash is a dev"

	split := strings.Split(input, " ")
	profaneWords := []string{"kerfuffle", "sharbert", "fornax"}

	for i, word := range split {
		replacement := strings.Repeat("*", 4)
		lowerCase := strings.ToLower(word)

		for _, profaneWord := range profaneWords {
			if lowerCase == profaneWord {
				split[i] = replacement
			}
		}
	}

	return strings.Join(split, " ")
}
