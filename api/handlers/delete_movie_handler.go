package handlers

import (
	"encoding/json"
	"net/http"
	"server/services"

	"github.com/gorilla/mux"
)

func DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		http.Error(w, "Id Parameter Required", http.StatusBadRequest)
		return
	}
	validId := services.FindId(movies, id)
	if !validId {
		http.Error(w, "There is no Movie with this ID", http.StatusBadRequest)
		return
	}
	services.DeleteMovieById(movies, id)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"Message": "Movie Deleted"})
}
