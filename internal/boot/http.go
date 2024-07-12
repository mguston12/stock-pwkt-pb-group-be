package boot

import (
	"log"
	"net/http"

	"skeleton/internal/config"

	"github.com/jmoiron/sqlx"

	skeletonData "skeleton/internal/data/skeleton"
	skeletonServer "skeleton/internal/delivery/http"
	skeletonHandler "skeleton/internal/delivery/http/skeleton"
	skeletonService "skeleton/internal/service/skeleton"
)

// HTTP will load configuration, do dependency injection and then start the HTTP server
func HTTP() error {

	err := config.Init()
	if err != nil {
		log.Fatalf("[CONFIG] Failed to initialize config: %v", err)
	}
	cfg := config.Get()
	// Open MySQL DB Connection
	db, err := sqlx.Open("mysql", cfg.Database.Master)
	if err != nil {
		log.Fatalf("[DB] Failed to initialize database connection: %v", err)
	}

	// Diganti dengan domain yang anda buat
	sd := skeletonData.New(db)
	ss := skeletonService.New(sd)
	sh := skeletonHandler.New(ss)

	server := skeletonServer.Server{
		Skeleton: sh,
	}

	if err := server.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}
