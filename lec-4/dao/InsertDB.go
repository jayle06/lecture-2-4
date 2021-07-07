package dao

import (
	"ocg.com/test/database"
	"ocg.com/test/model"
)

func InsertToDB(filmList []model.Film){
	data := database.Connection()
	defer data.Close()

	for _, film := range filmList {
		_, err := data.Query("INSERT INTO film VALUES ( ?, ?, ?, ?)", film.ID, film.Poster, film.Title, film.Rating)
		if err != nil {
			panic(err)
		}
	}
}
