package db

import "fmt"

func SongNotExists(id int) error {
	return fmt.Errorf("no song with id %d", id)
}

func InvalidLimit(limit int) error {
	return fmt.Errorf("invalid limit %d", limit)
}

func InvalidOffset(offset int) error {
	return fmt.Errorf("invalid offset %d", offset)
}
