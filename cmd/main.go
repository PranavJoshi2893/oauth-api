package main

import (
	"log"

	"github.com/PranavJoshi2893/oauth-api/internal/config"
	"github.com/PranavJoshi2893/oauth-api/internal/database"
	"github.com/PranavJoshi2893/oauth-api/internal/server"
)

func main() {
	cfg := config.Load()

	db, err := database.NewPostgres(cfg)

	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	server := server.New(cfg, db)

	if err := server.RunWithGracefulShutdown(); err != nil {
		log.Fatal(err)
	}

}
