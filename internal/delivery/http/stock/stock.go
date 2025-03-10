package stock

import (
	"context"
	"stock/internal/entity/stock"
)

type StockSvc interface {
	GetAllSpareparts(ctx context.Context) ([]stock.Sparepart, error)
	CreateSparepart(ctx context.Context, sparepart stock.Sparepart) error
	UpdateSparepart(ctx context.Context, sparepart stock.Sparepart) error
	DeleteSparepart(ctx context.Context, id string) error
}

type Handler struct {
	stockSvc StockSvc
}

func New(is StockSvc) *Handler {
	return &Handler{
		stockSvc: is,
	}
}
