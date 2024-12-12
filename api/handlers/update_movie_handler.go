package handlers

import (
	"encoding/json"
	"net/http"
	"server/models"
	"server/services"

	"github.com/gorilla/mux"
)

func UpdateMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	validId := services.FindId(movies, id)
	if !validId {
		http.Error(w, "There is no Movie with this ID", http.StatusBadRequest)
		return
	}
	services.DeleteMovieById(movies, id)
	var movieInformation models.Movie
	err := json.NewDecoder(r.Body).Decode(&movieInformation)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	movieInformation.ID = id
	movies = append(movies, movieInformation)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}
