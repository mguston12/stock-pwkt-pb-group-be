package stock

import (
	"context"
	"fmt"
	"stock-pwt/internal/entity/stock"

	"stock-pwt/pkg/errors"

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

func (d Data) GetAllSparepartsPage(ctx context.Context, keyword string, offset, limit int) ([]stock.Sparepart, error) {
	spareparts := []stock.Sparepart{}

	_keyword := "%" + keyword + "%"

	rows, err := d.stmt[getAllSparepartsPage].QueryxContext(ctx, _keyword, _keyword, _keyword, _keyword, offset, limit)
	if err != nil {
		return spareparts, errors.Wrap(err, "[DATA][GetAllSparepartsPage]")
	}

	defer rows.Close()

	for rows.Next() {
		sparepart := stock.Sparepart{}
		if err = rows.StructScan(&sparepart); err != nil {
			return spareparts, errors.Wrap(err, "[DATA][GetAllSparepartsPage]")
		}
		spareparts = append(spareparts, sparepart)
	}

	return spareparts, nil
}

func (d Data) GetAllSparepartsCount(ctx context.Context, keyword string) ([]stock.Sparepart, int, error) {
	spareparts := []stock.Sparepart{}
	var count int

	_keyword := "%" + keyword + "%"

	if err := d.stmt[getAllSparepartsCount].QueryRowxContext(ctx, _keyword, _keyword, _keyword, _keyword).Scan(&count); err != nil {
		return spareparts, count, errors.Wrap(err, "[DATA][GetAllSparepartsCount]")
	}

	return spareparts, count, nil
}

func (d Data) GetSparepartByID(ctx context.Context, id string) (stock.Sparepart, error) {
	sparepart := stock.Sparepart{}

	if err := d.stmt[getSparepartByID].QueryRowxContext(ctx, id).StructScan(&sparepart); err != nil {
		return sparepart, errors.Wrap(err, "[DATA][GetSparepartByID]")
	}

	return sparepart, nil
}

func (d Data) CreateSparepart(ctx context.Context, sparepart stock.Sparepart) error {
	_, err := d.stmt[createSparepart].ExecContext(ctx,
		sparepart.ID,
		sparepart.Nama,
		sparepart.Quantity,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][CreateSparepart]")
	}
	return nil
}

func (d Data) UpdateSparepart(ctx context.Context, sparepart stock.Sparepart) error {
	_, err := d.stmt[updateSparepart].ExecContext(ctx,
		sparepart.Nama,
		sparepart.Quantity,
		sparepart.ID,
	)

	fmt.Println("sparepart update", sparepart)

	if err != nil {
		return errors.Wrap(err, "[DATA][UpdateSparepart]")
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
