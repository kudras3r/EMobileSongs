package main

import (
	"github.com/kudras3r/EMobile/internal/config"
	"github.com/kudras3r/EMobile/internal/logger"
)

func main() {
	// TODO:

	// config
	config := config.Load()

	// logger
	log := logger.New(config.LogLevel)
	_ = log

	// router (chi / default ?)
	// run
}
