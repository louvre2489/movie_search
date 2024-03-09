package entity

import "time"

type (
	MovieID int64
)

type Movie struct {
	ID       MovieID   `json:"id" db:"id"`
	Title    string    `json:"title" db:"title"`
	Created  time.Time `json:"created" db:"created"`
	Modified time.Time `json:"modified" db:"modified"`
}

type Movies []*Movie
