package stock

import (
	"context"
	"stock/internal/entity/stock"

	"stock/pkg/errors"

	"github.com/jmoiron/sqlx"
)

func (d Data) GetAllCustomers(ctx context.Context) ([]stock.Customer, error) {
	var ( 
		rows  *sqlx.Rows
		datas []stock.Customer
		err   error
	)

	rows, err = d.stmt[getAllCustomers].QueryxContext(ctx)
	if err != nil {
		return datas, errors.Wrap(err, "[DATA][GetAllCustomers]")
	}

	for rows.Next() {
		var data stock.Customer
		err := rows.StructScan(&data)
		if err != nil {
			return datas, errors.Wrap(err, "[DATA][GetAllCustomers]")
		}
		datas = append(datas, data)
	}
	defer rows.Close()

	return datas, nil
}

func (d Data) CreateCustomer(ctx context.Context, customer stock.Customer) error {
	_, err := d.stmt[createCustomer].ExecContext(ctx,
		customer.ID,
		customer.Company,
		customer.NamaCustomer,
		customer.Alamat,
		customer.PIC,
		customer.UpdatedBy,
		customer.UpdatedAt,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][CreateContractDetail]")
	}
	return nil
}

func (d Data) UpdateCustomer(ctx context.Context, customer stock.Customer) error {
	_, err := d.stmt[updateCustomer].ExecContext(ctx,
		customer.NamaCustomer,
		customer.Alamat,
		customer.PIC,
		customer.UpdatedBy,
		customer.UpdatedAt,
		customer.ID,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][UpdateCustomer]")
	}
	return nil
}

func (d Data) DeleteCustomer(ctx context.Context, id string) error {
	_, err := d.stmt[deleteCustomer].ExecContext(ctx, id)

	if err != nil {
		return errors.Wrap(err, "[DATA][DeleteCustomer]")
	}
	return nil
}
