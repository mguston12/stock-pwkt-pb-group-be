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
	sparepart.HandleFunc("/create", s.Stock.CreateSparepart).Methods("POST")
	sparepart.HandleFunc("/update", s.Stock.UpdateSparepart).Methods("PUT")
	sparepart.HandleFunc("/delete", s.Stock.DeleteSparepart).Methods("DELETE")
	sparepart.HandleFunc("", s.Stock.GetSparepartsFiltered).Methods("GET")

	teknisi := r.PathPrefix("/teknisi").Subrouter()
	teknisi.HandleFunc("", s.Stock.GetAllTeknisi).Methods("GET")
	teknisi.HandleFunc("/detail", s.Stock.GetTeknisiByID).Methods("GET")
	teknisi.HandleFunc("/create", s.Stock.CreateTeknisi).Methods("POST")
	teknisi.HandleFunc("/update", s.Stock.UpdateTeknisi).Methods("PUT")
	teknisi.HandleFunc("/delete", s.Stock.DeleteTeknisi).Methods("DELETE")

	machine := r.PathPrefix("/machines").Subrouter()
	machine.HandleFunc("", s.Stock.GetAllMachines).Methods("GET")
	machine.HandleFunc("/customer", s.Stock.GetMachineByIDCustomer).Methods("GET")
	machine.HandleFunc("/detail", s.Stock.GetMachineByID).Methods("GET")
	machine.HandleFunc("/create", s.Stock.CreateMachine).Methods("POST")
	machine.HandleFunc("/update", s.Stock.UpdateMachine).Methods("PUT")
	machine.HandleFunc("/delete", s.Stock.DeleteMachine).Methods("DELETE")

	customer := r.PathPrefix("/customers").Subrouter()
	customer.HandleFunc("", s.Stock.GetAllCustomers).Methods("GET")
	customer.HandleFunc("/create", s.Stock.CreateCustomer).Methods("POST")
	customer.HandleFunc("/update", s.Stock.UpdateCustomer).Methods("PUT")
	customer.HandleFunc("/delete", s.Stock.DeleteCustomer).Methods("DELETE")

	request := r.PathPrefix("/requests").Subrouter()
	request.HandleFunc("", s.Stock.GetRequestsPagination).Methods("GET")
	request.HandleFunc("/create", s.Stock.CreateRequest).Methods("POST")
	request.HandleFunc("/update", s.Stock.UpdateRequest).Methods("PUT")
	request.HandleFunc("/delete", s.Stock.DeleteRequest).Methods("DELETE")

	user := r.PathPrefix("/users").Subrouter()
	user.HandleFunc("", s.Stock.GetUserByUsername).Methods("GET")
	user.HandleFunc("/create", s.Stock.CreateUser).Methods("POST")
	user.HandleFunc("/update", s.Stock.UpdateUser).Methods("PUT")
	user.HandleFunc("/delete", s.Stock.DeleteUser).Methods("DELETE")
	user.HandleFunc("/login", s.Stock.MatchPassword).Methods("POST")

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
