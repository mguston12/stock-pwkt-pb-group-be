package http

import (
	"net/http"

	"skeleton/pkg/grace"

	"github.com/rs/cors"
)

// SkeletonHandler ...
type SkeletonHandler interface {
	// Masukkan fungsi handler di sini
	GetSkeleton(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	Skeleton SkeletonHandler
}

// Serve is serving HTTP gracefully on port x ...
func (s *Server) Serve(port string) error {
	handler := cors.AllowAll().Handler(s.Handler())
	return grace.Serve(port, handler)
}
