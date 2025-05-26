package stock

import (
	"context"
	"math"
	"stock/internal/entity/stock"
	"stock/pkg/errors"
)

// GetAllCustomers ...
func (s Service) GetAllCustomers(ctx context.Context) ([]stock.Customer, error) {
	customer, err := s.data.GetAllCustomers(ctx)
	if err != nil {
		return customer, errors.Wrap(err, "[SERVICE][GetAllCustomers]")
	}

	return customer, nil
}

func (s Service) GetCustomersFiltered(ctx context.Context, keyword string, page, length int) ([]stock.Customer, int, error) {
	limit := length
	offset := (page - 1) * length
	var lastPage int

	if page != 0 && length != 0 {
		customers, count, err := s.data.GetAllCustomersCount(ctx, keyword)
		if err != nil {
			return customers, lastPage, errors.Wrap(err, "[SERVICE][GetCustomersFiltered][COUNT]")

		}

		lastPage = int(math.Ceil(float64(count) / float64(length)))

		customers, err = s.data.GetAllCustomersPage(ctx, keyword, offset, limit)
		if err != nil {
			return customers, lastPage, errors.Wrap(err, "[SERVICE][GetCustomersFiltered]")
		}
	}

	customers, err := s.data.GetAllCustomers(ctx)
	if err != nil {
		return customers, lastPage, errors.Wrap(err, "[SERVICE][GetCustomersFilteredf]")
	}

	return customers, lastPage, nil
}

func (s Service) CreateCustomer(ctx context.Context, customer stock.Customer) error {
	err := s.data.CreateCustomer(ctx, customer)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][CreateCustomer]")
	}

	return nil
}

func (s Service) UpdateCustomer(ctx context.Context, customer stock.Customer) error {
	err := s.data.UpdateCustomer(ctx, customer)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][UpdateCustomer]")
	}

	return nil
}

func (s Service) DeleteCustomer(ctx context.Context, id string) error {
	err := s.data.DeleteCustomer(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][DeleteCustomer]")
	}

	return nil
}
