package stock

import (
	"log"
	"net/http"
	httpHelper "stock-pwt/internal/delivery/http"
	"stock-pwt/internal/entity/stock"
	"stock-pwt/pkg/response"
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

	idCustomer := r.FormValue("id_customer")
	year := getIntQueryParam(r, "tahun")
	month := getIntQueryParam(r, "bulan")
	machineID := r.FormValue("machine_id")

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
