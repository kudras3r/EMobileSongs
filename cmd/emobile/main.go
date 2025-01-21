package main

import (
	"github.com/kudras3r/EMobile/internal/config"
	"github.com/kudras3r/EMobile/internal/db/migrate"
	"github.com/kudras3r/EMobile/internal/db/pg"
	"github.com/kudras3r/EMobile/internal/logger"
)

func main() {
	// TODO:

	// config
	config := config.Load()

	// logger
	log := logger.New(config.LogLevel)
	log.Info("logger is up")

	// storage
	db, err := pg.New(config.DB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Info("db is up")

	// migrate
	if err := migrate.CreateSongsTable(db.DB); err != nil {
		log.Fatal(err)
	}
	// router (chi / default ?)
	// run
}
