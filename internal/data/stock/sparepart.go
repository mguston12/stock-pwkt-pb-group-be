package stock

import (
	"context"
	"stock/internal/entity/stock"

	"stock/pkg/errors"

	"github.com/jmoiron/sqlx"
)

func (d Data) GetAllSpareparts(ctx context.Context) ([]stock.Sparepart, error) {
	var (
		rows  *sqlx.Rows
		datas []stock.Sparepart
		err   error
	)

	rows, err = d.stmt[getAllSpareparts].QueryxContext(ctx)
	if err != nil {
		return datas, errors.Wrap(err, "[DATA][GetAllSpareparts]")
	}

	for rows.Next() {
		var data stock.Sparepart
		err := rows.StructScan(&data)
		if err != nil {
			return datas, errors.Wrap(err, "[DATA][GetAllSpareparts]")
		}
		datas = append(datas, data)
	}
	defer rows.Close()

	return datas, nil
}

func (d Data) CreateSparepart(ctx context.Context, sparepart stock.Sparepart) error {
	_, err := d.stmt[createSparepart].ExecContext(ctx,
		sparepart.ID,
		sparepart.Nama,
		sparepart.Quantity,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][CreateContractDetail]")
	}
	return nil
}

func (d Data) UpdateSparepart(ctx context.Context, sparepart stock.Sparepart) error {
	_, err := d.stmt[updateSparepart].ExecContext(ctx,
		sparepart.Nama,
		sparepart.Quantity,
		sparepart.ID,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][CreateContractDetail]")
	}
	return nil
}

func (d Data) DeleteSparepart(ctx context.Context, id string) error {
	_, err := d.stmt[deleteSparepart].ExecContext(ctx, id)

	if err != nil {
		return errors.Wrap(err, "[DATA][DeleteSparepart]")
	}
	return nil
}
