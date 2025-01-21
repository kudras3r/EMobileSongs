package pg

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/kudras3r/EMobile/internal/config"
	_ "github.com/lib/pq"
)

func New(config config.DB) (*sqlx.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Pass, config.Name)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("cannot connect to db")
	}
	return db, nil
}
