package stock

import (
	"context"
	"stock/internal/entity/stock"
)

// Data ...
// Masukkan function dari package data ke dalam interface ini
type Data interface {
	GetAllSpareparts(ctx context.Context) ([]stock.Sparepart, error)
	CreateSparepart(ctx context.Context, sparepart stock.Sparepart) error
	UpdateSparepart(ctx context.Context, sparepart stock.Sparepart) error
	DeleteSparepart(ctx context.Context, id string) error
}

// Service ...
// Tambahkan variable sesuai banyak data layer yang dibutuhkan
type Service struct {
	data Data
}

// New ...
// Tambahkan parameter sesuai banyak data layer yang dibutuhkan
func New(data Data) Service {
	// Assign variable dari parameter ke object
	return Service{
		data: data,
	}
}
