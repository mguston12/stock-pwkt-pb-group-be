package stock

import (
	"context"
	"stock-pwt/internal/entity/stock"
	"stock-pwt/pkg/errors"

	"github.com/jmoiron/sqlx"
)

func (d Data) GetAllMachineHistory(ctx context.Context, idMachine, idCustomer string, offset, limit int) ([]stock.MachineHistory, error) {
	var (
		rows  *sqlx.Rows
		datas []stock.MachineHistory
		err   error
	)

	rows, err = d.stmt[getAllMachineHistory].QueryxContext(ctx,
		"%"+idMachine+"%", idMachine,
		"%"+idCustomer+"%", idCustomer,
		offset, limit,
	)
	if err != nil {
		return datas, errors.Wrap(err, "[DATA][GetAllMachineHistory]")
	}
	defer rows.Close()

	for rows.Next() {
		var data stock.MachineHistory
		err := rows.StructScan(&data)
		if err != nil {
			return datas, errors.Wrap(err, "[DATA][GetAllMachineHistory][Scan]")
		}
		datas = append(datas, data)
	}

	return datas, nil
}

func (d Data) GetMachineHistoryByID(ctx context.Context, idMachine string) ([]stock.MachineHistory, error) {
	var (
		rows  *sqlx.Rows
		datas []stock.MachineHistory
		err   error
	)

	rows, err = d.stmt[getMachineHistoryByID].QueryxContext(ctx, idMachine)
	if err != nil {
		return datas, errors.Wrap(err, "[DATA][GetMachineHistoryByID]")
	}
	defer rows.Close()

	for rows.Next() {
		var data stock.MachineHistory
		err := rows.StructScan(&data)
		if err != nil {
			return datas, errors.Wrap(err, "[DATA][GetMachineHistoryByID][Scan]")
		}
		datas = append(datas, data)
	}

	return datas, nil
}

func (d Data) GetAllMachineHistoryCount(ctx context.Context, idMachine, idCustomer string) (int, error) {
	var count int

	err := d.stmt[getAllMachineHistoryCount].QueryRowContext(ctx, idMachine, idMachine, idCustomer, idCustomer).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "[DATA][GetAllMachineHistoryCount]")
	}

	return count, nil
}

func (d Data) GetLastMachineHistoryByMachineID(ctx context.Context, id string) (stock.MachineHistory, error) {
	machine_history := stock.MachineHistory{}

	if err := d.stmt[getLastMachineHistoryByMachineID].QueryRowxContext(ctx, id).StructScan(&machine_history); err != nil {
		return machine_history, errors.Wrap(err, "[DATA][GetLastMachineHistoryByMachineID]")
	}

	return machine_history, nil
}

func (d Data) CreateMachineHistory(ctx context.Context, history stock.MachineHistory) error {
	_, err := d.stmt[createMachineHistory].ExecContext(ctx,
		history.IDMachine,
		history.IDCustomer,
		history.TanggalMulai,
		history.Status,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][CreateMachineHistory]")
	}
	return nil
}

func (d Data) UpdateMachineHistory(ctx context.Context, history stock.MachineHistory) error {
	_, err := d.stmt[updateMachineHistory].ExecContext(ctx,
		history.IDCustomer,
		history.TanggalMulai,
		history.TanggalSelesai,
		history.Status,
		history.IDHistory,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][UpdateMachineHistory]")
	}
	return nil
}

func (d Data) DeleteMachineHistory(ctx context.Context, id int) error {
	_, err := d.stmt[deleteMachineHistory].ExecContext(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[DATA][DeleteMachineHistory]")
	}
	return nil
}
