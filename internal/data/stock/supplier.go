package stock

import (
	"context"
	"stock/internal/entity/stock"

	"stock/pkg/errors"

	"github.com/jmoiron/sqlx"
)

func (d Data) GetSuppliers(ctx context.Context) ([]stock.Supplier, error) {
	var (
		rows  *sqlx.Rows
		datas []stock.Supplier
		err   error
	)

	rows, err = d.stmt[getSuppliers].QueryxContext(ctx)
	if err != nil {
		return datas, errors.Wrap(err, "[DATA][GetSuppliers]")
	}

	for rows.Next() {
		var data stock.Supplier
		err := rows.StructScan(&data)
		if err != nil {
			return datas, errors.Wrap(err, "[DATA][GetSuppliers]")
		}
		datas = append(datas, data)
	}
	defer rows.Close()

	return datas, nil
}

func (d Data) GetSuppliersPage(ctx context.Context, nama string, offset, limit int) ([]stock.Supplier, error) {
	suppliers := []stock.Supplier{}

	_keyword := "%" + nama + "%"

	rows, err := d.stmt[getSuppliersPage].QueryxContext(ctx, _keyword, _keyword, offset, limit)
	if err != nil {
		return suppliers, errors.Wrap(err, "[DATA][GetSuppliersPage]")
	}

	defer rows.Close()

	for rows.Next() {
		supplier := stock.Supplier{}
		if err = rows.StructScan(&supplier); err != nil {
			return suppliers, errors.Wrap(err, "[DATA][GetSuppliersPage]")
		}
		suppliers = append(suppliers, supplier)
	}

	return suppliers, nil
}

func (d Data) GetSuppliersCount(ctx context.Context, nama string) ([]stock.Supplier, int, error) {
	suppliers := []stock.Supplier{}
	var count int
	_keyword := "%" + nama + "%"

	if err := d.stmt[getSuppliersCount].QueryRowxContext(ctx, _keyword, _keyword).Scan(&count); err != nil {
		return suppliers, count, errors.Wrap(err, "[DATA][GetSuppliersCount]")
	}

	return suppliers, count, nil
}

func (d Data) CreateSupplier(ctx context.Context, supplier stock.Supplier) error {
	_, err := d.stmt[createSupplier].ExecContext(ctx,
		supplier.Nama,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][CreateSupplier]")
	}
	return nil
}

func (d Data) UpdateSupplier(ctx context.Context, supplier stock.Supplier) error {
	_, err := d.stmt[updateSupplier].ExecContext(ctx,
		supplier.Nama,
		supplier.ID,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][UpdateSupplier]")
	}
	return nil
}

func (d Data) DeleteSupplier(ctx context.Context, id string) error {
	_, err := d.stmt[deleteSupplier].ExecContext(ctx, id)

	if err != nil {
		return errors.Wrap(err, "[DATA][DeleteSupplier]")
	}
	return nil
}
