package skeleton

import (
	"context"

	"github.com/pkg/errors"
)

// Data ...
// Masukkan function dari package data ke dalam interface ini
type Data interface {
	GetSkeleton(ctx context.Context) error
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

// GetSkeleton ...
func (s Service) GetSkeleton(ctx context.Context) error {
	err := s.data.GetSkeleton(ctx)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][GetSkeleton]")
	}
	return nil
}
