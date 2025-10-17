package stock

import (
	"context"
	"fmt"
	"math"
	"stock-pwt/internal/entity/stock"
	"stock-pwt/pkg/errors"
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

		return customers, lastPage, nil
	}

	customers, err := s.data.GetAllCustomers(ctx)
	if err != nil {
		return customers, lastPage, errors.Wrap(err, "[SERVICE][GetCustomersFiltered]")
	}

	return customers, lastPage, nil
}

func (s Service) GetCustomerID(ctx context.Context, company int) (string, error) {
	com := ""

	switch company {
	case 1:
		com = "PB"
	case 2:
		com = "PBM"
	case 3:
		com = "MMU"
	case 4:
		com = "PWT"
	}

	id, err := s.data.FetchAndIncreaseCounter(ctx, company)
	if err != nil {
		return "", errors.Wrap(err, "[SERVICE][GetCustomerFiltered]")
	}

	return fmt.Sprintf("%s-%06d", com, id), nil
}

func (s Service) CreateCustomer(ctx context.Context, customer stock.Customer) error {
	customerID, err := s.GetCustomerID(ctx, customer.Company)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][CreateCustomer][1]")
	}

	customer.ID = customerID

	err = s.data.CreateCustomer(ctx, customer)
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
