package model

type Film struct {
	ID     int64  `json:"id"`
	Poster string `json:"poster"`
	Title  string `json:"title"`
	Rating float64 `json:"rating"`
}