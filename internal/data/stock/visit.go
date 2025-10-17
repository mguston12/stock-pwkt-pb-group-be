package stock

import (
	"context"
	"log"
	"stock-pwt/internal/entity/stock"

	"stock-pwt/pkg/errors"

	"github.com/jmoiron/sqlx"
)

func (d Data) GetVisitsByID(ctx context.Context, id string) ([]stock.Visit, error) {
	var (
		rows  *sqlx.Rows
		datas []stock.Visit
		err   error
	)

	log.Println("id : ", id)

	rows, err = d.stmt[getVisits].QueryxContext(ctx, id)
	if err != nil {
		return datas, errors.Wrap(err, "[DATA][GetVisits]")
	}

	for rows.Next() {
		var data stock.Visit
		err := rows.StructScan(&data)
		if err != nil {
			return datas, errors.Wrap(err, "[DATA][GetVisits]")
		}
		datas = append(datas, data)
	}
	defer rows.Close()

	return datas, nil
}

func (d Data) CreateVisit(ctx context.Context, visit stock.Visit) error {
	_, err := d.stmt[createVisit].ExecContext(ctx,
		visit.MachineID,
		visit.Notes,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][CreateVisit]")
	}
	return nil
}

func (d Data) UpdateVisit(ctx context.Context, visit stock.Visit) error {
	_, err := d.stmt[updateVisit].ExecContext(ctx,
		visit.Notes,
		visit.ID,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][UpdateVisit]")
	}
	return nil
}

func (d Data) DeleteVisit(ctx context.Context, id int) error {
	log.Println(id)
	_, err := d.stmt[deleteVisit].ExecContext(ctx, id)

	if err != nil {
		return errors.Wrap(err, "[DATA][DeleteVisit]")
	}
	return nil
}
