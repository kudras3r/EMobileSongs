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
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Pass, config.Name)
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

func (s *Storage) fetchSongs(limit, offset int, filters map[string]string) ([]models.Song, error) {
	songs := make([]models.Song, limit)
	var b strings.Builder

	b.WriteString("SELECT * FROM songs WHERE TRUE")
	for k, v := range filters {
		b.WriteString(fmt.Sprintf(" AND \"%s\" = '%s'", k, v))
	}

	b.WriteString(fmt.Sprintf(" LIMIT %d OFFSET %d", limit, offset))

	if err := s.DB.Select(&songs, b.String()); err != nil {
		return nil, err
	}

	return songs, nil
}

func (s *Storage) GetSongs() ([]models.Song, error) {
	// example
	a := make(map[string]string)
	a["group_performer"] = "asd"
	sn, err := s.fetchSongs(2, 0, a)
	if err != nil {
		return nil, err
	}

	return sn, nil
}
