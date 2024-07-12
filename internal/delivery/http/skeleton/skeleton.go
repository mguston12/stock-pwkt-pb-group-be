package skeleton

import (
	"context"
	"log"
	"net/http"
	"skeleton/pkg/response"
)

type SkeletonSvc interface {
	GetSkeleton(ctx context.Context) error
}

type Handler struct {
	skeletonSvc SkeletonSvc
}

func New(is SkeletonSvc) *Handler {
	return &Handler{
		skeletonSvc: is,
	}
}

func (h *Handler) GetSkeleton(w http.ResponseWriter, r *http.Request) {
	resp := response.Response{}
	defer resp.RenderJSON(w, r)

	log.Printf("[INFO][%s][%s] %s", r.RemoteAddr, r.Method, r.URL)
}
