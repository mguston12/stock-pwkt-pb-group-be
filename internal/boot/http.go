package boot

import (
	"log"
	"net/http"

	"stock/internal/config"

	"github.com/jmoiron/sqlx"

	stockData "stock/internal/data/stock"
	stockServer "stock/internal/delivery/http"
	stockHandler "stock/internal/delivery/http/stock"
	stockService "stock/internal/service/stock"
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
	sd := stockData.New(db)
	ss := stockService.New(sd)
	sh := stockHandler.New(ss)

	server := stockServer.Server{
		Stock: sh,
	}

	if err := server.Serve(cfg.Server.Port); err != http.ErrServerClosed {
		return err
	}

	return nil
}
