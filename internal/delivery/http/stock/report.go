package stock

import (
	"log"
	"net/http"
	"strconv"
)

func (h *Handler) ExportExcel(w http.ResponseWriter, r *http.Request) {
	// Parse query parameters
	bulan, err := strconv.Atoi(r.FormValue("bulan"))
	if err != nil {
		http.Error(w, "Parameter 'bulan' tidak valid", http.StatusBadRequest)
		return
	}
	tahun, err := strconv.Atoi(r.FormValue("tahun"))
	if err != nil {
		http.Error(w, "Parameter 'tahun' tidak valid", http.StatusBadRequest)
		return
	}

	// Generate the Excel file
	fileBytes, fileName, err := h.stockSvc.ExportExcel(r.Context(), bulan, tahun)
	if err != nil {
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		http.Error(w, "Gagal membuat file Excel", http.StatusInternalServerError)
		return
	}

	// Set headers for file download
	w.Header().Set("Content-Disposition", "attachment; filename="+fileName)
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.WriteHeader(http.StatusOK)
	w.Write(fileBytes)

	log.Printf("[INFO][%s][%s] %s | Exported: %s", r.RemoteAddr, r.Method, r.URL, fileName)
}
