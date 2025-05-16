package stock

import (
	"context"
	"stock/internal/entity/stock"

	"stock/pkg/errors"

	"github.com/jmoiron/sqlx"
)

func (d Data) GetAllMachines(ctx context.Context) ([]stock.Machine, error) {
	var (
		rows  *sqlx.Rows
		datas []stock.Machine
		err   error
	)

	rows, err = d.stmt[getAllMachines].QueryxContext(ctx)
	if err != nil {
		return datas, errors.Wrap(err, "[DATA][GetAllMachines]")
	}

	for rows.Next() {
		var data stock.Machine
		err := rows.StructScan(&data)
		if err != nil {
			return datas, errors.Wrap(err, "[DATA][GetAllMachines]")
		}
		datas = append(datas, data)
	}
	defer rows.Close()

	return datas, nil
}

func (d Data) GetMachineByID(ctx context.Context, id string) (stock.Machine, error) {
	machine := stock.Machine{}

	if err := d.stmt[getMachineByID].QueryRowxContext(ctx, id).StructScan(&machine); err != nil {
		return machine, errors.Wrap(err, "[DATA][GetMachineByID]")
	}

	return machine, nil
}

func (d Data) GetMachineByIDCustomer(ctx context.Context, id string) ([]stock.Machine, error) {
	var (
		rows  *sqlx.Rows
		datas []stock.Machine
		err   error
	)

	rows, err = d.stmt[getMachineByIDCustomer].QueryxContext(ctx, id)
	if err != nil {
		return datas, errors.Wrap(err, "[DATA][GetMachineByIDCustomer]")
	}

	for rows.Next() {
		var data stock.Machine
		err := rows.StructScan(&data)
		if err != nil {
			return datas, errors.Wrap(err, "[DATA][GetMachineByIDCustomer]")
		}
		datas = append(datas, data)
	}
	defer rows.Close()

	return datas, nil
}

func (d Data) CreateMachine(ctx context.Context, machine stock.Machine) error {
	_, err := d.stmt[createMachine].ExecContext(ctx,
		machine.ID,
		machine.Type,
		machine.Customer,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][CreateContractDetail]")
	}
	return nil
}

func (d Data) UpdateMachine(ctx context.Context, machine stock.Machine) error {
	_, err := d.stmt[updateMachine].ExecContext(ctx,
		machine.Customer,
		machine.ID,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][UpdateMachine]")
	}
	return nil
}

func (d Data) DeleteMachine(ctx context.Context, id string) error {
	_, err := d.stmt[deleteMachine].ExecContext(ctx, id)

	if err != nil {
		return errors.Wrap(err, "[DATA][DeleteMachine]")
	}
	return nil
}
