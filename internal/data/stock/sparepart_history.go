package stock

import (
	"context"
	"stock-pwt/internal/entity/stock"

	"stock-pwt/pkg/errors"

	"github.com/jmoiron/sqlx"
)

func (d Data) GetAllSparepartHistory(ctx context.Context, id_teknisi, id_mesin, id_sparepart string, offset, limit int) ([]stock.SparepartHistory, error) {
	var (
		rows  *sqlx.Rows
		datas []stock.SparepartHistory
		err   error
	)

	rows, err = d.stmt[getAllSparepartHistory].QueryxContext(ctx, id_teknisi, id_teknisi, id_mesin, id_mesin, id_sparepart, id_sparepart, offset, limit)
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

func (d Data) GetAllSparepartHistoryCount(ctx context.Context, id_teknisi, id_mesin, id_sparepart string) ([]stock.SparepartHistory, int, error) {
	history := []stock.SparepartHistory{}
	var count int

	if err := d.stmt[getAllSparepartHistoryCount].QueryRowxContext(ctx, id_teknisi, id_teknisi, id_mesin, id_mesin, id_sparepart, id_sparepart).Scan(&count); err != nil {
		return history, count, errors.Wrap(err, "[DATA][GetAllSparepartHistoryCount]")
	}

	return history, count, nil
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
		history.Quantity,
		history.Counter,
		history.CounterColour,
		history.CounterColourA3,
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
		history.Quantity,
		history.Counter,
		history.CounterColour,
		history.CounterColourA3,
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
