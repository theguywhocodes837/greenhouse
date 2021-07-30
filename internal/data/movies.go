package data

import "time"

type Movie struct {
	ID       int64     `json:"id"`
	CreateAt time.Time `json:"created_at"`
	Title    string    `json:"title"`
	Year     int32     `json:"year"`
	Runtime  Runtime   `json:"runtime"`
	Generes  []string  `json:"generes"`
	Version  int32     `json:"version"`
}
