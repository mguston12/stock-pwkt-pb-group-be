package stock

import (
	"log"
	"net/http"
	httpHelper "stock/internal/delivery/http"
	"stock/internal/entity/stock"
	"stock/pkg/response"
	"strconv"
)

func (h *Handler) GetSparepartCost(w http.ResponseWriter, r *http.Request) {
	var (
		result   interface{}
		metadata interface{}
		err      error
		resp     response.Response
	)
	defer resp.RenderJSON(w, r)

	// Menangani parameter query
	idCustomer := r.FormValue("id_customer")
	year := getIntQueryParam(r, "year")
	month := getIntQueryParam(r, "month")
	machineID := r.FormValue("machine_id") // Add this to get the machine_id from query

	// Set filter with the parameters
	var filter stock.SparepartCostFilter
	if idCustomer != "" {
		filter.CustomerID = &idCustomer
	}
	filter.Year = year
	filter.Month = month

	if machineID != "" {
		filter.MachineID = &machineID
	}

	// Memanggil service untuk mendapatkan data cost sparepart
	ctx := r.Context()
	result, err = h.stockSvc.GetSparepartCost(ctx, filter)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error()) // Parse error and render it in JSON response
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	// Add metadata (if any, or leave it empty)
	resp.Data = result
	resp.Metadata = metadata

	// Log the success info
	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

// Fungsi untuk mengambil parameter string sebagai pointer *string
func getStringQueryParam(r *http.Request, key string) *string {
	val := r.FormValue(key)
	if val == "" {
		return nil // Kembalikan nil jika parameter tidak ada
	}
	return &val // Kembalikan pointer ke string
}

func getIntQueryParam(r *http.Request, key string) *int {
	val := r.FormValue(key)
	if val == "" {
		return nil // Kembalikan nil jika tidak ada parameter
	}
	result, err := strconv.Atoi(val)
	if err != nil {
		return nil // Kembalikan nil jika parsing gagal
	}
	return &result // Kembalikan pointer ke result
}

// Fungsi tambahan untuk memeriksa apakah string kosong dan mengembalikan *string yang nil jika kosong
func nilIfEmpty(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}
