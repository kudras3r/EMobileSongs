package models

import "time"

type Song struct {
	ID          int       `db:"id" json:"id"`
	Song        string    `db:"title" json:"song"`
	Group       string    `db:"group_performer" json:"group"`
	Link        string    `db:"link" json:"link"`
	Text        string    `db:"lyrics" json:"text"`
	ReleaseDate time.Time `db:"release_date" json:"releaseDate"`
}
