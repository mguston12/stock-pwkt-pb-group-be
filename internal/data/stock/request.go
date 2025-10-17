package stock

import (
	"context"
	"database/sql"
	"log"
	"stock-pwt/internal/entity/stock"

	"stock-pwt/pkg/errors"

	"github.com/jmoiron/sqlx"
)

func (d Data) GetAllRequests(ctx context.Context) ([]stock.Request, error) {
	var (
		rows  *sqlx.Rows
		datas []stock.Request
		err   error
	)

	rows, err = d.stmt[getAllRequests].QueryxContext(ctx)
	if err != nil {
		return datas, errors.Wrap(err, "[DATA][GetAllRequests]")
	}

	for rows.Next() {
		var data stock.Request
		err := rows.StructScan(&data)
		if err != nil {
			return datas, errors.Wrap(err, "[DATA][GetAllRequests]")
		}
		datas = append(datas, data)
	}
	defer rows.Close()

	return datas, nil
}

func (d Data) GetRequestsPage(ctx context.Context, teknisi, status string, offset, limit int) ([]stock.Request, error) {
	requests := []stock.Request{}

	_keyword := "%" + teknisi + "%"

	rows, err := d.stmt[getRequestsPage].QueryxContext(ctx, _keyword, _keyword, _keyword, _keyword, status, status, offset, limit)
	if err != nil {
		return requests, errors.Wrap(err, "[DATA][GetRequestsPage]")
	}

	defer rows.Close()

	for rows.Next() {
		request := stock.Request{}
		if err = rows.StructScan(&request); err != nil {
			return requests, errors.Wrap(err, "[DATA][GetRequestsPage]")
		}
		requests = append(requests, request)
	}

	return requests, nil
}

func (d Data) GetRequestsCount(ctx context.Context, teknisi, status string) ([]stock.Request, int, error) {
	requests := []stock.Request{}
	var count int

	_keyword := "%" + teknisi + "%"

	if err := d.stmt[getRequestsCount].QueryRowxContext(ctx, _keyword, _keyword, _keyword, _keyword, status, status).Scan(&count); err != nil {
		return requests, count, errors.Wrap(err, "[DATA][GetRequestsCount]")
	}

	return requests, count, nil
}

func (d Data) GetRequestByID(ctx context.Context, id int) (stock.Request, error) {
	request := stock.Request{}

	if err := d.stmt[getRequestByID].QueryRowxContext(ctx, id).StructScan(&request); err != nil {
		return request, errors.Wrap(err, "[DATA][GetRequestByID]")
	}

	return request, nil
}

func (d Data) CreateRequest(ctx context.Context, request stock.Request) error {
	_, err := d.stmt[createRequest].ExecContext(ctx,
		request.Teknisi,
		request.Mesin,
		request.Sparepart,
		request.Quantity,
		request.Status,
		request.UpdatedBy,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][CreateRequest]")
	}
	return nil
}

func (d Data) UpdateRequest(ctx context.Context, request stock.Request) error {
	_, err := d.stmt[updateRequest].ExecContext(ctx,
		request.Mesin,
		request.Sparepart,
		request.Quantity,
		request.Status,
		request.UpdatedBy,
		request.ID,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][UpdateRequest]")
	}
	return nil
}

func (d Data) DeleteRequest(ctx context.Context, id string) error {
	_, err := d.stmt[deleteRequest].ExecContext(ctx, id)

	if err != nil {
		return errors.Wrap(err, "[DATA][DeleteRequest]")
	}
	return nil
}

func (d Data) DeleteRequestByIDSparepart(ctx context.Context, id string) error {
	result, err := d.stmt[deleteRequestByIDSparepart].ExecContext(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[DATA][DeleteRequestByIDSparepart]")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "[DATA][DeleteRequestByIDSparepart][RowsAffected]")
	}

	if rowsAffected == 0 {
		log.Printf("[INFO] Tidak ada pembelian sparepart dengan id %s untuk dihapus", id)
	}

	return nil
}

func (d Data) CheckSparepartValidOrNot(ctx context.Context, id string) (int, error) {
	var count int

	query := `SELECT COUNT(*) from stock.request WHERE id_sparepart = ? AND status_request = "Disetujui";`

	if err := d.db.QueryRowxContext(ctx, query, id).Scan(&count); err != nil {
		return count, errors.Wrap(err, "[DATA][CheckSparepartValidOrNot]")
	}

	return count, nil
}

func (d Data) BeginTx(ctx context.Context) (*sql.Tx, error) {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "[DATA][BeginTx]")
	}
	return tx, nil
}

func (d Data) CreateRequestTx(ctx context.Context, tx *sql.Tx, request stock.Request) error {
	const query = `
		INSERT INTO request (
			id_teknisi, id_mesin, id_sparepart, quantity, status_request, updated_by
		) VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := tx.ExecContext(ctx, query,
		request.Teknisi,
		request.Mesin,
		request.Sparepart,
		request.Quantity,
		request.Status,
		request.UpdatedBy,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][CreateRequestTx]")
	}

	return nil
}
