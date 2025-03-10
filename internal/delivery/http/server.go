package http

import (
	"net/http"

	"stock/pkg/grace"

	"github.com/rs/cors"
)


type StockHandler interface {
	// Masukkan fungsi handler di sini
	GetAllSparepart(w http.ResponseWriter, r *http.Request)
	CreateSparepart(w http.ResponseWriter, r *http.Request)
	UpdateSparepart(w http.ResponseWriter, r *http.Request)
	DeleteSparepart(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	Stock StockHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	handler := cors.AllowAll().Handler(s.Handler())
	return grace.Serve(port, handler)
}
