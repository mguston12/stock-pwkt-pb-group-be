package stock

import (
	"context"
	"stock-pwt/internal/entity/stock"

	"stock-pwt/pkg/errors"

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

func (d Data) GetAllCustomersPage(ctx context.Context, keyword string, offset, limit int) ([]stock.Customer, error) {
	customers := []stock.Customer{}

	_keyword := "%" + keyword + "%"

	rows, err := d.stmt[getAllCustomersPage].QueryxContext(ctx, _keyword, _keyword, _keyword, _keyword, offset, limit)
	if err != nil {
		return customers, errors.Wrap(err, "[DATA][GetAllCustomersPage]")
	}

	defer rows.Close()

	for rows.Next() {
		customer := stock.Customer{}
		if err = rows.StructScan(&customer); err != nil {
			return customers, errors.Wrap(err, "[DATA][GetAllCustomersPage]")
		}
		customers = append(customers, customer)
	}

	return customers, nil
}

func (d Data) GetAllCustomersCount(ctx context.Context, keyword string) ([]stock.Customer, int, error) {
	customers := []stock.Customer{}
	var count int

	_keyword := "%" + keyword + "%"

	if err := d.stmt[getAllCustomersCount].QueryRowxContext(ctx, _keyword, _keyword, _keyword, _keyword).Scan(&count); err != nil {
		return customers, count, errors.Wrap(err, "[DATA][GetAllCustomersCount]")
	}

	return customers, count, nil
}

func (d Data) CreateCustomer(ctx context.Context, customer stock.Customer) error {
	_, err := d.stmt[createCustomer].ExecContext(ctx,
		customer.ID,
		customer.Company,
		customer.NamaCustomer,
		customer.Alamat,
		customer.UpdatedBy,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][CreateContractDetail]")
	}
	return nil
}

func (d Data) UpdateCustomer(ctx context.Context, customer stock.Customer) error {
	_, err := d.stmt[updateCustomer].ExecContext(ctx,
		customer.Alamat,
		customer.NamaCustomer,
		customer.UpdatedBy,
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

func (d Data) FetchAndIncreaseCounter(ctx context.Context, company int) (int, error) {
	counter := 0
	tx, err := d.db.Beginx()
	if err != nil {
		return counter, errors.Wrap(err, "[DATA][FetchAndIncreaseCounter]")
	}
	if err := tx.QueryRowxContext(ctx, `SELECT count FROM counter WHERE company = ?`, company).Scan(&counter); err != nil {
		return counter, errors.Wrap(err, "[DATA][FetchAndIncreaseCounter][A]")
	}
	if _, err := tx.ExecContext(ctx, `UPDATE counter SET count = count + 1 WHERE company = ?`, company); err != nil {
		return counter, errors.Wrap(err, "[DATA][FetchAndIncreaseCounter][B]")
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return counter, errors.Wrap(err, "[DATA][FetchAndIncreaseCounter]")
	}
	return counter, nil
}
