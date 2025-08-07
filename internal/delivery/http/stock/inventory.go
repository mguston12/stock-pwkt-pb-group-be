package stock

import (
	"encoding/json"
	"log"
	"net/http"
	httpHelper "stock/internal/delivery/http"
	"stock/internal/entity/stock"
	"stock/pkg/response"
)

func (h *Handler) GetAllInventory(w http.ResponseWriter, r *http.Request) {
	var (
		result   interface{}
		metadata interface{}
		err      error
		resp     response.Response
	)
	defer resp.RenderJSON(w, r)

	ctx := r.Context()
	result, err = h.stockSvc.GetAllInventory(ctx)

	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	resp.Data = result
	resp.Metadata = metadata

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) GetInventoryByID(w http.ResponseWriter, r *http.Request) {
	var (
		result   interface{}
		metadata interface{}
		err      error
		resp     response.Response
	)
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	result, err = h.stockSvc.GetInventoryByID(ctx, r.FormValue("id"))
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	resp.Data = result
	resp.Metadata = metadata

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) InventoryUsage(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	input := stock.InventoryUsage{}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	err = h.stockSvc.InventoryUsage(ctx, input)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) InventoryUsageBatch(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	var inputs []stock.InventoryUsage // harus slice/array di sini

	err := json.NewDecoder(r.Body).Decode(&inputs)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	err = h.stockSvc.InventoryUsageBatch(ctx, inputs)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	log.Printf("[INFO][%s][%s] %s - Batch inventory usage success", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) CreateInventory(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	inventory := stock.Inventory{}

	err := json.NewDecoder(r.Body).Decode(&inventory)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	err = h.stockSvc.CreateInventory(ctx, inventory)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) UpdateInventory(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	inventory := stock.Inventory{}

	err := json.NewDecoder(r.Body).Decode(&inventory)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	err = h.stockSvc.UpdateInventory(ctx, inventory)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) DeleteInventory(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	err := h.stockSvc.DeleteInventory(ctx, r.FormValue("teknisi"), r.FormValue("sparepart"))
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}
