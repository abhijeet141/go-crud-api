package services

import "server/models"

func DeleteMovieById(movies []models.Movie, id string) {
	for index, value := range movies {
		if value.ID == id {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
}
func FindId(movies []models.Movie, id string) bool {
	for _, value := range movies {
		if value.ID == id {
			return true
		}
	}
	return false
}
func FindMovieById(movies []models.Movie, id string) *models.Movie {
	for _, value := range movies {
		if value.ID == id {
			return &value
		}
	}
	return nil
}
