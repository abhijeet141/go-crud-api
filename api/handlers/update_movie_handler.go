package handlers

import (
	"encoding/json"
	"net/http"
	"server/models"
	"strconv"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gorilla/mux"
)

func UpdateMovieHandler(w http.ResponseWriter, r *http.Request) {
	o := orm.NewOrm()
	vars := mux.Vars(r)
	movieId := vars["id"]
	var updateMovie models.Movie
	err := json.NewDecoder(r.Body).Decode(&updateMovie)
	if err != nil {
		http.Error(w, "Invalid JSON data", http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(movieId)
	if err != nil {
		http.Error(w, "Invalid Movie ID", http.StatusBadRequest)
		return
	}
	movie := models.Movie{ID: id}
	err = o.Read(&movie)
	if err == orm.ErrNoRows {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}
	if updateMovie.Isbn != "" {
		movie.Isbn = updateMovie.Isbn
	}
	if updateMovie.Title != "" {
		movie.Title = updateMovie.Title
	}
	if updateMovie.Director.FirstName != "" {
		movie.Director.FirstName = updateMovie.Director.FirstName
	}
	if updateMovie.Director.LastName != "" {
		movie.Director.LastName = updateMovie.Director.LastName
	}
	_, err = o.Update(&movie)
	if err != nil {
		http.Error(w, "Failed to update the movie", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}
