package stock

import (
	"context"
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
