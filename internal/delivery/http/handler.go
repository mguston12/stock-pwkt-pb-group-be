package http

import (
	"errors"
	"log"
	"net/http"

	"stock/pkg/response"

	"github.com/gorilla/mux"
)

// Handler will initialize mux router and register handler
func (s *Server) Handler() *mux.Router {
	r := mux.NewRouter()
	// Jika tidak ditemukan, jangan diubah.
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	// Tambahan Prefix di depan API endpoint
	router := r.PathPrefix("").Subrouter()
	// Health Check
	router.HandleFunc("", defaultHandler).Methods("GET")
	router.HandleFunc("/", defaultHandler).Methods("GET")
	// Routes

	sparepart := r.PathPrefix("/spareparts").Subrouter()
	sparepart.HandleFunc("", s.Stock.GetAllSparepart).Methods("GET")
	sparepart.HandleFunc("/create", s.Stock.CreateSparepart).Methods("POST")
	sparepart.HandleFunc("/update", s.Stock.UpdateSparepart).Methods("PUT")
	sparepart.HandleFunc("/delete", s.Stock.DeleteSparepart).Methods("DELETE")

	return r
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Go Skeleton API"))
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	var (
		resp   *response.Response
		err    error
		errRes response.Error
	)
	resp = &response.Response{}
	defer resp.RenderJSON(w, r)

	err = errors.New("404 Not Found")

	if err != nil {
		// Error response handling
		errRes = response.Error{
			Code:   404,
			Msg:    "404 Not Found",
			Status: true,
		}

		log.Printf("[ERROR] %s %s - %v\n", r.Method, r.URL, err)
		resp.StatusCode = 404
		resp.Error = errRes
		return
	}
}
