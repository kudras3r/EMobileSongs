package songs

import "github.com/kudras3r/EMobile/internal/models"

var FilterFields []string = []string{"song", "group", "link", "releaseDate", "versesCount", "text"}

type Service interface {
	GetSongs(limit, offset int, filters *map[string]string) ([]models.Song, error)
	GetSongText(id, limit, offset int) ([]models.Verse, error)
	UpdateSong(id int, updated models.Song) (int, error)
	DeleteSong(id int) (int, error)
}

func ConvertFField(fField string) string {
	var res string
	switch fField {
	case "song":
		res = "title"
	case "group":
		res = "group_performer"
	case "releaseDate":
		res = "release_date"
	case "versesCount":
		res = "verses_count"
	default:
		res = fField
	}

	return res
}
