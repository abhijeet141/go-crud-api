package handlers

import (
	"encoding/json"
	"net/http"
	"server/services"

	"github.com/gorilla/mux"
)

func GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	movie := services.FindMovieById(movies, id)
	if movie == nil {
		http.Error(w, "There is no Movie with this ID", http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}
