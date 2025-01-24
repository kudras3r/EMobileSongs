package models

import "time"

type Song struct {
	ID          int       `db:"id" json:"id"`
	Song        string    `db:"title" json:"song"`
	Group       string    `db:"group_performer" json:"group"`
	Link        string    `db:"link" json:"link"`
	ReleaseDate time.Time `db:"release_date" json:"releaseDate"`
	VersesCount int       `db:"verses_count" json:"versesCount"`
	Text        string    `json:"text"`
}

type Verse struct {
	ID      int    `db:"id" json:"id"`
	Song_id int    `db:"song_id" json:"song_id"`
	Number  int    `db:"num" json:"number"`
	Text    string `db:"lyrics" json:"text"`
}
