package stock

import (
	"encoding/json"
	"log"
	"net/http"
	httpHelper "stock-pwt/internal/delivery/http"
	"stock-pwt/internal/entity/stock"
	"stock-pwt/pkg/response"
)

// GetAllReturnInventory ...
func (h *Handler) GetAllReturnInventory(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	result, err := h.stockSvc.GetAllReturnInventory(ctx)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	resp.Data = result
	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

// GetReturnInventoryByStatus ...
func (h *Handler) GetReturnInventoryByStatus(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()
	status := r.FormValue("status")

	result, err := h.stockSvc.GetReturnInventoryByStatus(ctx, status)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	resp.Data = result
	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

// ProcessReturnSparepart ...
func (h *Handler) ProcessReturnSparepart(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()
	var input stock.ReturnInventory

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		resp = httpHelper.ParseErrorCode("invalid request body")
		log.Printf("[ERROR][%s][%s] %s | Reason: %v", r.RemoteAddr, r.Method, r.URL, err)
		return
	}

	if err := h.stockSvc.ProcessReturnSparepart(ctx, input); err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %v", r.RemoteAddr, r.Method, r.URL, err)
		return
	}

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

// ApproveReturnInventory ...
func (h *Handler) ApproveReturnInventory(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()
	var input stock.ReturnInventory

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		resp = httpHelper.ParseErrorCode("invalid request body")
		log.Printf("[ERROR][%s][%s] %s | Reason: %v", r.RemoteAddr, r.Method, r.URL, err)
		return
	}

	if err := h.stockSvc.ApproveReturnInventory(ctx, input); err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %v", r.RemoteAddr, r.Method, r.URL, err)
		return
	}

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}
