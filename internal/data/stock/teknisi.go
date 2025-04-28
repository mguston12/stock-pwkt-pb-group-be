package stock

import (
	"context"
	"stock/internal/entity/stock"

	"stock/pkg/errors"

	"github.com/jmoiron/sqlx"
)

func (d Data) GetAllTeknisi(ctx context.Context) ([]stock.Teknisi, error) {
	var (
		rows  *sqlx.Rows
		datas []stock.Teknisi
		err   error
	)

	rows, err = d.stmt[getAllTeknisi].QueryxContext(ctx)
	if err != nil {
		return datas, errors.Wrap(err, "[DATA][GetAllTeknisis]")
	}

	for rows.Next() {
		var data stock.Teknisi
		err := rows.StructScan(&data)
		if err != nil {
			return datas, errors.Wrap(err, "[DATA][GetAllTeknisis]")
		}
		datas = append(datas, data)
	}
	defer rows.Close()

	return datas, nil
}

func (d Data) GetTeknisiByID(ctx context.Context, id string) (stock.Teknisi, error) {
	teknisi := stock.Teknisi{}

	if err := d.stmt[getTeknisiByID].QueryRowxContext(ctx, id).StructScan(&teknisi); err != nil {
		return teknisi, errors.Wrap(err, "[DATA][GetTeknisiByID]")
	}

	return teknisi, nil
}

func (d Data) CreateTeknisi(ctx context.Context, teknisi stock.Teknisi) error {
	_, err := d.stmt[createTeknisi].ExecContext(ctx,
		teknisi.ID,
		teknisi.Nama,
		teknisi.ActiveYN,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][CreateTeknisi]")
	}
	return nil
}

func (d Data) UpdateTeknisi(ctx context.Context, teknisi stock.Teknisi) error {
	_, err := d.stmt[updateTeknisi].ExecContext(ctx,
		teknisi.Nama,
		teknisi.ActiveYN,
		teknisi.ID,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][UpdateTeknisi]")
	}
	return nil
}

func (d Data) DeleteTeknisi(ctx context.Context, id string) error {
	_, err := d.stmt[deleteTeknisi].ExecContext(ctx, id)

	if err != nil {
		return errors.Wrap(err, "[DATA][DeleteTeknisi]")
	}
	return nil
}
