package handlers

import (
	"encoding/json"
	"net/http"
	"server/models"

	"github.com/beego/beego/v2/client/orm"
)

// var movies = []models.Movie{
// 	{
// 		ID:    "1",
// 		Isbn:  "123456",
// 		Title: "Inception",
// 		Director: &models.Director{
// 			FirstName: "Christopher",
// 			LastName:  "Nolan",
// 		},
// 	},
// 	{
// 		ID:    "2",
// 		Isbn:  "789012",
// 		Title: "The Matrix",
// 		Director: &models.Director{
// 			FirstName: "Lana",
// 			LastName:  "Wachowski",
// 		},
// 	},
// 	{
// 		ID:    "3",
// 		Isbn:  "345678",
// 		Title: "Interstellar",
// 		Director: &models.Director{
// 			FirstName: "Christopher",
// 			LastName:  "Nolan",
// 		},
// 	},
// 	{
// 		ID:    "4",
// 		Isbn:  "901234",
// 		Title: "The Social Network",
// 		Director: &models.Director{
// 			FirstName: "David",
// 			LastName:  "Fincher",
// 		},
// 	},
// 	{
// 		ID:    "5",
// 		Isbn:  "567890",
// 		Title: "Pulp Fiction",
// 		Director: &models.Director{
// 			FirstName: "Quentin",
// 			LastName:  "Tarantino",
// 		},
// 	},
// 	{
// 		ID:    "6",
// 		Isbn:  "512043",
// 		Title: "Avengers",
// 		Director: &models.Director{
// 			FirstName: "Alan",
// 			LastName:  "Root",
// 		},
// 	},
// }

func GetMoviesHandler(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie
	o := orm.NewOrm()
	_, err := o.QueryTable("movie").All(&movies)
	if err != nil {
		http.Error(w, "Error retrieving movies", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}
