package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/beego/beego/v2/client/orm"
	"github.com/gorilla/mux"
)

func DeleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	o := orm.NewOrm()
	vars := mux.Vars(r)
	id := vars["id"]
	num, err := o.QueryTable("movie").Filter("id", id).Delete()
	if err != nil {
		http.Error(w, "Failed to delete the movie", http.StatusInternalServerError)
		return
	}
	if num == 0 {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"Message": "Movie Deleted Successfully"})
}
