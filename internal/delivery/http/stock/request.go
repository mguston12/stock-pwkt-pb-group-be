package stock

import (
	"encoding/json"
	"log"
	"net/http"
	httpHelper "stock/internal/delivery/http"
	"stock/internal/entity/stock"
	"stock/pkg/response"
	"strconv"
)

func (h *Handler) GetAllRequests(w http.ResponseWriter, r *http.Request) {
	var (
		result   interface{}
		metadata interface{}
		err      error
		resp     response.Response
	)
	defer resp.RenderJSON(w, r)

	ctx := r.Context()
	result, err = h.stockSvc.GetAllRequests(ctx)

	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	resp.Data = result
	resp.Metadata = metadata

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) GetRequestsPagination(w http.ResponseWriter, r *http.Request) {
	var (
		result   interface{}
		metadata interface{}
		err      error
		resp     response.Response
	)
	defer resp.RenderJSON(w, r)

	page, _ := strconv.Atoi(r.FormValue("page"))
	length, _ := strconv.Atoi(r.FormValue("length"))

	ctx := r.Context()
	result, metadata, err = h.stockSvc.GetRequestsPagination(ctx, r.FormValue("keyword"), r.FormValue("status"), page, length)

	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	resp.Data = result
	resp.Metadata = metadata

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) CreateRequest(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	request := stock.Request{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	err = h.stockSvc.CreateRequest(ctx, request)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) CreateRequests(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	var requests []stock.Request

	err := json.NewDecoder(r.Body).Decode(&requests)
	if err != nil {
		resp = httpHelper.ParseErrorCode("invalid_request_payload")
		log.Printf("[ERROR][%s][%s] %s | JSON Decode Failed: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	if len(requests) == 0 {
		resp = httpHelper.ParseErrorCode("empty_request_list")
		log.Printf("[ERROR][%s][%s] %s | Reason: Empty request list", r.RemoteAddr, r.Method, r.URL)
		return
	}

	err = h.stockSvc.CreateRequests(ctx, requests)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Failed to create requests: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	log.Printf("[INFO][%s][%s] %s | Successfully created %d requests", r.RemoteAddr, r.Method, r.URL, len(requests))
}

func (h *Handler) UpdateRequest(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	request := stock.Request{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	err = h.stockSvc.UpdateRequest(ctx, request)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) CancelRequest(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	request := stock.Request{}

	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	err = h.stockSvc.CancelRequest(ctx, request)
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}

func (h *Handler) DeleteRequest(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	ctx := r.Context()

	err := h.stockSvc.DeleteRequest(ctx, r.FormValue("id"))
	if err != nil {
		resp = httpHelper.ParseErrorCode(err.Error())
		log.Printf("[ERROR][%s][%s] %s | Reason: %s", r.RemoteAddr, r.Method, r.URL, err.Error())
		return
	}

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}
