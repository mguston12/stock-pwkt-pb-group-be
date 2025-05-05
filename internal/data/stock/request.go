package stock

import (
	"context"
	"stock/internal/entity/stock"

	"stock/pkg/errors"

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

	rows, err := d.stmt[getRequestsPage].QueryxContext(ctx, teknisi, teknisi, status, status, offset, limit)
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

	if err := d.stmt[getRequestsCount].QueryRowxContext(ctx, teknisi, teknisi, status, status).Scan(&count); err != nil {
		return requests, count, errors.Wrap(err, "[DATA][GetRequestsCount]")
	}

	return requests, count, nil
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
