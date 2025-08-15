package http

import (
	"errors"
	"log"
	"net/http"

	"stock/internal/middleware"
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
	sparepart.Use(middleware.AuthMiddleware)
	sparepart.HandleFunc("/create", s.Stock.CreateSparepart).Methods("POST")
	sparepart.HandleFunc("/update", s.Stock.UpdateSparepart).Methods("PUT")
	sparepart.HandleFunc("/delete", s.Stock.DeleteSparepart).Methods("DELETE")
	sparepart.HandleFunc("", s.Stock.GetSparepartsFiltered).Methods("GET")
	sparepart.HandleFunc("/cost", s.Stock.GetSparepartCost).Methods("GET")

	teknisi := r.PathPrefix("/teknisi").Subrouter()
	teknisi.Use(middleware.AuthMiddleware)
	teknisi.HandleFunc("", s.Stock.GetAllTeknisi).Methods("GET")
	teknisi.HandleFunc("/detail", s.Stock.GetTeknisiByID).Methods("GET")
	teknisi.HandleFunc("/create", s.Stock.CreateTeknisi).Methods("POST")
	teknisi.HandleFunc("/update", s.Stock.UpdateTeknisi).Methods("PUT")
	teknisi.HandleFunc("/delete", s.Stock.DeleteTeknisi).Methods("DELETE")

	machine := r.PathPrefix("/machines").Subrouter()
	machine.Use(middleware.AuthMiddleware)
	machine.HandleFunc("", s.Stock.GetMachinesFiltered).Methods("GET")
	machine.HandleFunc("/customer", s.Stock.GetMachineByIDCustomer).Methods("GET")
	machine.HandleFunc("/detail", s.Stock.GetMachineByID).Methods("GET")
	machine.HandleFunc("/create", s.Stock.CreateMachine).Methods("POST")
	machine.HandleFunc("/replace", s.Stock.ReplaceCustomerMachine).Methods("POST")
	machine.HandleFunc("/deactivate", s.Stock.DeactivateMachine).Methods("POST")
	machine.HandleFunc("/update", s.Stock.UpdateMachine).Methods("PUT")
	machine.HandleFunc("/delete", s.Stock.DeleteMachine).Methods("DELETE")

	customer := r.PathPrefix("/customers").Subrouter()
	customer.Use(middleware.AuthMiddleware)
	customer.HandleFunc("", s.Stock.GetCustomersFiltered).Methods("GET")
	customer.HandleFunc("/create", s.Stock.CreateCustomer).Methods("POST")
	customer.HandleFunc("/update", s.Stock.UpdateCustomer).Methods("PUT")
	customer.HandleFunc("/delete", s.Stock.DeleteCustomer).Methods("DELETE")

	request := r.PathPrefix("/requests").Subrouter()
	request.Use(middleware.AuthMiddleware)
	request.HandleFunc("", s.Stock.GetRequestsPagination).Methods("GET")
	request.HandleFunc("/create", s.Stock.CreateRequest).Methods("POST")
	request.HandleFunc("/batch", s.Stock.CreateRequests).Methods("POST")
	request.HandleFunc("/update", s.Stock.UpdateRequest).Methods("PUT")
	request.HandleFunc("/delete", s.Stock.DeleteRequest).Methods("DELETE")

	// Public route: /users/login
	router.HandleFunc("/users/login", s.Stock.MatchPassword).Methods("POST")
	router.HandleFunc("/users/update", s.Stock.UpdateUser).Methods("PUT")
	router.HandleFunc("/croncustomers", s.Stock.ImportCustomersFromExcel).Methods("GET")
	router.HandleFunc("/cronmachines", s.Stock.ImportMachineFromExcel).Methods("GET")

	visit := router.PathPrefix("/visits").Subrouter()
	visit.HandleFunc("", s.Stock.GetVisitsByID).Methods("GET")
	visit.HandleFunc("/create", s.Stock.CreateVisit).Methods("POST")
	visit.HandleFunc("/update", s.Stock.UpdateVisit).Methods("PUT")
	visit.HandleFunc("/delete", s.Stock.DeleteVisit).Methods("DELETE")

	// Protected routes: /users/*
	user := router.PathPrefix("/users").Subrouter()
	user.Use(middleware.AuthMiddleware)
	user.HandleFunc("", s.Stock.GetUserByUsername).Methods("GET")
	user.HandleFunc("/create", s.Stock.CreateUser).Methods("POST")
	user.HandleFunc("/delete", s.Stock.DeleteUser).Methods("DELETE")

	inventory := r.PathPrefix("/inventory").Subrouter()
	inventory.Use(middleware.AuthMiddleware)
	inventory.HandleFunc("", s.Stock.GetAllInventory).Methods("GET")
	inventory.HandleFunc("/detail", s.Stock.GetInventoryByID).Methods("GET")
	inventory.HandleFunc("/usage-batch", s.Stock.InventoryUsageBatch).Methods("POST")
	inventory.HandleFunc("/usage", s.Stock.InventoryUsage).Methods("POST")
	inventory.HandleFunc("/create", s.Stock.CreateInventory).Methods("POST")
	inventory.HandleFunc("/update", s.Stock.UpdateInventory).Methods("PUT")
	inventory.HandleFunc("/delete", s.Stock.DeleteInventory).Methods("DELETE")

	pembeliansp := r.PathPrefix("/purchase").Subrouter()
	pembeliansp.Use(middleware.AuthMiddleware)
	pembeliansp.HandleFunc("", s.Stock.GetPembelianSparepart).Methods("GET")
	pembeliansp.HandleFunc("/{id}", s.Stock.GetPembelianSparepartByID).Methods("GET")
	pembeliansp.HandleFunc("/supplier/{id_supplier}", s.Stock.GetPembelianSparepartBySupplier).Methods("GET")
	pembeliansp.HandleFunc("/create", s.Stock.CreatePembelianSparepart).Methods("POST")
	pembeliansp.HandleFunc("/update", s.Stock.UpdatePembelianSparepart).Methods("PUT")
	pembeliansp.HandleFunc("/delete", s.Stock.DeletePembelianSparepart).Methods("DELETE")

	history := r.PathPrefix("/histories").Subrouter()
	history.Use(middleware.AuthMiddleware)
	history.HandleFunc("", s.Stock.GetAllSparepartHistory).Methods("GET")
	history.HandleFunc("/", s.Stock.GetAllSparepartHistory).Methods("GET")
	history.HandleFunc("/report", s.Stock.ExportExcel).Methods("GET")

	suppliers := r.PathPrefix("/suppliers").Subrouter()
	suppliers.Use(middleware.AuthMiddleware)
	suppliers.HandleFunc("", s.Stock.GetSuppliersPagination).Methods("GET")
	suppliers.HandleFunc("/create", s.Stock.CreateSupplier).Methods("POST")
	suppliers.HandleFunc("/update", s.Stock.UpdateSupplier).Methods("PUT")
	suppliers.HandleFunc("/delete", s.Stock.DeleteSupplier).Methods("DELETE")

	returnInv := r.PathPrefix("/return-inventory").Subrouter()
	returnInv.Use(middleware.AuthMiddleware)
	returnInv.HandleFunc("", s.Stock.GetAllReturnInventory).Methods("GET")
	returnInv.HandleFunc("/status", s.Stock.GetReturnInventoryByStatus).Methods("GET")
	returnInv.HandleFunc("/return", s.Stock.ProcessReturnSparepart).Methods("POST")
	returnInv.HandleFunc("/approve", s.Stock.ApproveReturnInventory).Methods("POST")

	machine_history := r.PathPrefix("/machine-history").Subrouter()
	machine_history.Use(middleware.AuthMiddleware)
	machine_history.HandleFunc("", s.Stock.GetAllMachineHistories).Methods("GET")
	machine_history.HandleFunc("/detail", s.Stock.GetMachineHistoryByID).Methods("GET")
	machine_history.HandleFunc("/create", s.Stock.CreateMachineHistory).Methods("POST")
	machine_history.HandleFunc("/update", s.Stock.UpdateMachineHistory).Methods("PUT")
	machine_history.HandleFunc("/delete", s.Stock.DeleteMachineHistory).Methods("DELETE")

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
