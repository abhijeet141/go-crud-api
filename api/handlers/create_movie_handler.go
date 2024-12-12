package handlers

import (
	"encoding/json"
	"math/rand/v2"
	"net/http"
	"server/models"
	"strconv"
)

func CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movieInformation models.Movie
	err := json.NewDecoder(r.Body).Decode(&movieInformation)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	movieInformation.ID = strconv.Itoa(rand.IntN(100000000))
	movies = append(movies, movieInformation)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}
