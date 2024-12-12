package routes

import (
	"server/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/movies", handlers.GetMoviesHandler).Methods("GET")
	r.HandleFunc("/movies/{id}", handlers.GetMovieHandler).Methods("GET")
	r.HandleFunc("/movies", handlers.CreateMovieHandler).Methods("POST")
	r.HandleFunc("/movies/{id}", handlers.UpdateMovieHandler).Methods("PUT")
	r.HandleFunc("/movies/{id}", handlers.DeleteMovieHandler).Methods("DELETE")
	return r
}
