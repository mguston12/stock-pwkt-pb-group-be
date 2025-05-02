package stock

import (
	"context"
	"stock/internal/entity/stock"

	"stock/pkg/errors"

	"github.com/jmoiron/sqlx"
)

func (d Data) GetAllSparepartHistory(ctx context.Context) ([]stock.SparepartHistory, error) {
	var (
		rows  *sqlx.Rows
		datas []stock.SparepartHistory
		err   error
	)

	rows, err = d.stmt[getAllSparepartHistory].QueryxContext(ctx)
	if err != nil {
		return datas, errors.Wrap(err, "[DATA][GetAllSparepartHistory]")
	}

	for rows.Next() {
		var data stock.SparepartHistory
		err := rows.StructScan(&data)
		if err != nil {
			return datas, errors.Wrap(err, "[DATA][GetAllSpareparts]")
		}
		datas = append(datas, data)
	}
	defer rows.Close()

	return datas, nil
}

func (d Data) GetSparepartHistoryByID(ctx context.Context, id string) ([]stock.SparepartHistory, error) {
	var (
		rows  *sqlx.Rows
		datas []stock.SparepartHistory
		err   error
	)

	rows, err = d.stmt[getSparepartHistoryByID].QueryxContext(ctx, id)
	if err != nil {
		return datas, errors.Wrap(err, "[DATA][GetSparepartHistoryByID]")
	}

	for rows.Next() {
		var data stock.SparepartHistory
		err := rows.StructScan(&data)
		if err != nil {
			return datas, errors.Wrap(err, "[DATA][GetSparepartHistoryByID]")
		}
		datas = append(datas, data)
	}
	defer rows.Close()

	return datas, nil
}

func (d Data) CreateSparepartHistory(ctx context.Context, history stock.SparepartHistory) error {
	_, err := d.stmt[createSparepartHistory].ExecContext(ctx,
		history.IDHistory,
		history.IDTeknisi,
		history.IDMachine,
		history.IDSparepart,
		history.IDRequest,
		history.Quantity,
		history.UpdatedBy,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][CreateSparepartHistory]")
	}
	return nil
}

func (d Data) UpdateSparepartHistory(ctx context.Context, history stock.SparepartHistory) error {
	_, err := d.stmt[updateSparepartHistory].ExecContext(ctx,
		history.IDTeknisi,
		history.IDMachine,
		history.IDSparepart,
		history.IDRequest,
		history.Quantity,
		history.UpdatedBy,
		history.UpdatedAt,
		history.IDHistory,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][UpdateSparepartHistory]")
	}
	return nil
}

func (d Data) DeleteSparepartHistory(ctx context.Context, id string) error {
	_, err := d.stmt[deleteSparepartHistory].ExecContext(ctx, id)

	if err != nil {
		return errors.Wrap(err, "[DATA][DeleteSparepartHistory]")
	}
	return nil
}
