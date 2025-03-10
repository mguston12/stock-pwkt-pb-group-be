package stock

import (
	"encoding/json"
	"log"
	"net/http"
	httpHelper "stock/internal/delivery/http"
	"stock/internal/entity/stock"
	"stock/pkg/response"
)

func (h *Handler) GetAllSparepart(w http.ResponseWriter, r *http.Request) {
	var (
		result   interface{}
		metadata interface{}
		err      error
		resp     response.Response
	)
	defer resp.RenderJSON(w, r)

	ctx := r.Context()
	result, err = h.stockSvc.GetAllSpareparts(ctx)

	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	resp.Data = result
	resp.Metadata = metadata

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) CreateSparepart(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	sparepart := stock.Sparepart{}

	err := json.NewDecoder(r.Body).Decode(&sparepart)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	err = h.stockSvc.CreateSparepart(ctx, sparepart)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) UpdateSparepart(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	customer := stock.Sparepart{}

	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	err = h.stockSvc.UpdateSparepart(ctx, customer)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) DeleteSparepart(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	err := h.stockSvc.DeleteSparepart(ctx, r.FormValue("id"))
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}
