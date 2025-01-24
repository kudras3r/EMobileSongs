package pg

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/kudras3r/EMobile/internal/config"
	"github.com/kudras3r/EMobile/internal/models"
	_ "github.com/lib/pq"
)

type Storage struct {
	DB *sqlx.DB
}

func New(config config.DB) (*Storage, error) {
	connStr := fmt.Sprintf(
		`host=%s port=%d user=%s 
		password=%s dbname=%s sslmode=disable`,
		config.Host, config.Port, config.User,
		config.Pass, config.Name)

	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &Storage{DB: db}, nil
}

func (s *Storage) CloseConnection() {
	s.DB.Close()
}

func (s *Storage) GetConnection() *sql.DB {
	return s.DB.DB
}

func (s *Storage) GetSongs(limit, offset int, filters *map[string]string) ([]models.Song, error) {
	/* for paginating we use offset-based query */

	songs := make([]models.Song, limit)
	var b strings.Builder

	// take songs info
	b.WriteString("SELECT * FROM songs WHERE TRUE")
	if filters != nil {
		for k, v := range *filters {
			b.WriteString(fmt.Sprintf(" AND \"%s\" = '%s'", k, v))
		}
	}

	b.WriteString(fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset))

	if err := s.DB.Select(&songs, b.String()); err != nil {
		return nil, err
	}

	// take songs texts
	for i := 0; i < len(songs); i++ {
		verses := make([]string, songs[i].VersesCount)
		query := fmt.Sprintf(
			`SELECT lyrics 
			 FROM verses 
			 WHERE song_id = %d 
			 ORDER BY num ASC`,
			songs[i].ID)

		if err := s.DB.Select(&verses, query); err != nil {
			return nil, err
		}
		b.Reset()
		for _, v := range verses {
			b.WriteString(v)
		}
		songs[i].Text = b.String()
	}

	return songs, nil
}

func (s *Storage) GetSongText(id, limit, offset int) ([]models.Verse, error) {
	verses := make([]models.Verse, limit)
	query := fmt.Sprintf(
		`SELECT * 
		 FROM verses 
		 WHERE song_id = %d 
		 ORDER BY num ASC 
		 LIMIT %d OFFSET %d`,
		id, limit, offset)

	if err := s.DB.Select(&verses, query); err != nil {
		return nil, err
	}

	return verses, nil
}

func (s *Storage) DeleteSong(id int) (int, error) {
	query := fmt.Sprintf("DELETE FROM songs WHERE id = %d", id)
	res, err := s.DB.Exec(query)
	if err != nil {
		return -1, err
	}

	rAffected, err := res.RowsAffected()
	if err != nil {
		return -1, err
	}
	if rAffected == 0 {
		return -1, fmt.Errorf("no song found with id: %d", id)
	}

	return id, nil
}
