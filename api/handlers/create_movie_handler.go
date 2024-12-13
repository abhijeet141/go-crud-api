package handlers

import (
	"encoding/json"
	"net/http"
	"server/models"

	"github.com/beego/beego/v2/client/orm"
)

func CreateMovieHandler(w http.ResponseWriter, r *http.Request) {
	var movieInformation models.Movie
	err := json.NewDecoder(r.Body).Decode(&movieInformation)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}
	o := orm.NewOrm()
	director := &models.Director{
		FirstName: movieInformation.Director.FirstName,
		LastName:  movieInformation.Director.LastName,
	}
	_, err = o.Insert(director)
	if err != nil {
		http.Error(w, "Failed to create director", http.StatusInternalServerError)
		return
	}
	_, err = o.Insert(&movieInformation)
	if err != nil {
		http.Error(w, "Failed to create book", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Movie created successfully"})
}
