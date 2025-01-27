package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/kudras3r/EMobile/internal/api"
	"github.com/kudras3r/EMobile/internal/config"
	"github.com/kudras3r/EMobile/internal/db/migrate"
	"github.com/kudras3r/EMobile/internal/db/pg"
	"github.com/kudras3r/EMobile/pkg/logger"
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

	// router
	r := chi.NewRouter()
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)

	api.RegisterRoutes(r, log, storage, config)

	server := &http.Server{
		Addr:         config.Server.Address,
		Handler:      r,
		ReadTimeout:  config.Server.RWTimeout,
		WriteTimeout: config.Server.RWTimeout,
		IdleTimeout:  config.Server.IdleTimeout,
	}

	// run
	log.Info(fmt.Sprintf("server is running on %s", config.Server.Address))
	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

	log.Info("gracefully shut down")
}
