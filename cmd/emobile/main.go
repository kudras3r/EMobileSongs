package main

import (
	"fmt"

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
	storage, err := pg.New(config.DB)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Info("database is up")
	}
	defer storage.CloseConnection()

	// migrate
	if err := migrate.CreateSongsTable(storage.GetConnection()); err != nil {
		log.Fatal(err)
	} else {
		log.Info("migration applied")
	}

	// example
	s, err := storage.GetSongs()
	if err != nil {
		log.Error(err)
	}
	for _, v := range s {
		fmt.Println(v)
	}
	// router (chi / default ?)
	// run
}
