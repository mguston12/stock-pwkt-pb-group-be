package http

import (
	"net/http"

	"stock/pkg/grace"

	"github.com/rs/cors"
)

type StockHandler interface {
	// Masukkan fungsi handler di sini
	GetAllSparepart(w http.ResponseWriter, r *http.Request)
	GetSparepartsFiltered(w http.ResponseWriter, r *http.Request)
	CreateSparepart(w http.ResponseWriter, r *http.Request)
	UpdateSparepart(w http.ResponseWriter, r *http.Request)
	DeleteSparepart(w http.ResponseWriter, r *http.Request)

	GetAllTeknisi(w http.ResponseWriter, r *http.Request)
	GetTeknisiByID(w http.ResponseWriter, r *http.Request)
	CreateTeknisi(w http.ResponseWriter, r *http.Request)
	UpdateTeknisi(w http.ResponseWriter, r *http.Request)
	DeleteTeknisi(w http.ResponseWriter, r *http.Request)

	GetAllMachines(w http.ResponseWriter, r *http.Request)
	GetMachineByID(w http.ResponseWriter, r *http.Request)
	GetMachineByIDCustomer(w http.ResponseWriter, r *http.Request)
	CreateMachine(w http.ResponseWriter, r *http.Request)
	UpdateMachine(w http.ResponseWriter, r *http.Request)
	DeleteMachine(w http.ResponseWriter, r *http.Request)

	GetAllCustomers(w http.ResponseWriter, r *http.Request)
	CreateCustomer(w http.ResponseWriter, r *http.Request)
	UpdateCustomer(w http.ResponseWriter, r *http.Request)
	DeleteCustomer(w http.ResponseWriter, r *http.Request)

	GetAllRequests(w http.ResponseWriter, r *http.Request)
	GetRequestsPagination(w http.ResponseWriter, r *http.Request)
	CreateRequest(w http.ResponseWriter, r *http.Request)
	UpdateRequest(w http.ResponseWriter, r *http.Request)
	DeleteRequest(w http.ResponseWriter, r *http.Request)

	GetUserByUsername(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
	UpdateUser(w http.ResponseWriter, r *http.Request)
	DeleteUser(w http.ResponseWriter, r *http.Request)

	MatchPassword(w http.ResponseWriter, r *http.Request)

	GetAllInventory(w http.ResponseWriter, r *http.Request)
	GetInventoryByID(w http.ResponseWriter, r *http.Request)
	CreateInventory(w http.ResponseWriter, r *http.Request)
	UpdateInventory(w http.ResponseWriter, r *http.Request)
	DeleteInventory(w http.ResponseWriter, r *http.Request)
}

// Server ...
type Server struct {
	Stock StockHandler
}

// Serve is serving HTTP gracefully on port x ...

func (s *Server) Serve(port string) error {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"*",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(s.Handler())

	return grace.Serve(port, handler)
}
