package boot

import (
	"log"
	"net/http"

	"stock-pwt/internal/config"

	"github.com/jmoiron/sqlx"

	stockData "stock-pwt/internal/data/stock"
	stockServer "stock-pwt/internal/delivery/http"
	stockHandler "stock-pwt/internal/delivery/http/stock"
	stockService "stock-pwt/internal/service/stock"
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

	_, err = db.Exec("SET time_zone = '+07:00'")
	if err != nil {
		log.Fatalf("[DB] Failed to set time zone: %v", err)
		return err
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
