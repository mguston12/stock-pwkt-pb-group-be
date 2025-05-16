package stock

import (
	"context"
	"math"
	"stock/internal/entity/stock"
	"stock/pkg/errors"
)

// GetSuppliers ...
func (s Service) GetSuppliers(ctx context.Context) ([]stock.Supplier, error) {
	supplier, err := s.data.GetSuppliers(ctx)
	if err != nil {
		return supplier, errors.Wrap(err, "[SERVICE][GetSuppliers]")
	}

	return supplier, nil
}

func (s Service) GetSuppliersPagination(ctx context.Context, keyword string, page, length int) ([]stock.Supplier, int, error) {
	limit := length
	offset := (page - 1) * length
	var lastPage int

	if page != 0 && length != 0 {
		suppliers, count, err := s.data.GetSuppliersCount(ctx, keyword)
		if err != nil {
			return suppliers, lastPage, errors.Wrap(err, "[SERVICE][GetSuppliersFiltered][COUNT]")

		}

		lastPage = int(math.Ceil(float64(count) / float64(length)))

		suppliers, err = s.data.GetSuppliersPage(ctx, keyword, offset, limit)
		if err != nil {
			return suppliers, lastPage, errors.Wrap(err, "[SERVICE][GetSuppliersFiltered]")
		}

		return suppliers, lastPage, nil
	}

	suppliers, err := s.data.GetSuppliers(ctx)
	if err != nil {
		return suppliers, lastPage, errors.Wrap(err, "[SERVICE][GetSuppliersFiltered]")
	}

	return suppliers, lastPage, nil
}

func (s Service) CreateSupplier(ctx context.Context, supplier stock.Supplier) error {
	err := s.data.CreateSupplier(ctx, supplier)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][CreateSupplier]")
	}

	return nil
}

func (s Service) UpdateSupplier(ctx context.Context, supplier stock.Supplier) error {
	err := s.data.UpdateSupplier(ctx, supplier)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][UpdateSupplier]")
	}

	return nil
}

func (s Service) DeleteSupplier(ctx context.Context, id string) error {
	err := s.data.DeleteSupplier(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][DeleteSupplier]")
	}

	return nil
}
