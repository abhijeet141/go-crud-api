package handlers

import (
	"encoding/json"
	"net/http"
	"server/models"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gorilla/mux"
)

func GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	o := orm.NewOrm()
	vars := mux.Vars(r)
	id := vars["id"]
	err := o.QueryTable("movie").Filter("id", id).One(&movie)
	if err == orm.ErrNoRows {
		http.Error(w, "Movie Not Found", http.StatusNotFound)
	} else if err != nil {
		http.Error(w, "Failed to fetch movie", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movie)
}
