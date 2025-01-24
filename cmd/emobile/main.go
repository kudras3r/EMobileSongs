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
	m := make(map[string]string)
	m["title"] = "asd_song"
	s, err := storage.GetSongs(2, 0, &m)
	if err != nil {
		log.Error(err)
	}
	for _, v := range s {
		fmt.Println(v)
	}
	v, err := storage.GetSongText(1, 2, 1)
	for _, k := range v {
		fmt.Println(k)
	}

	// id, err := storage.DeleteSong(2)
	// if err != nil {
	// 	log.Error(err)
	// }
	// fmt.Println(id)
	// router (chi / default ?)
	// run
}
