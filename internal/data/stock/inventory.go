package stock

import (
	"context"
	"stock/internal/entity/stock"

	"stock/pkg/errors"

	"github.com/jmoiron/sqlx"
)

func (d Data) GetAllInventory(ctx context.Context) ([]stock.Inventory, error) {
	var (
		rows  *sqlx.Rows
		datas []stock.Inventory
		err   error
	)

	rows, err = d.stmt[getAllInventory].QueryxContext(ctx)
	if err != nil {
		return datas, errors.Wrap(err, "[DATA][GetAllInventory]")
	}

	for rows.Next() {
		var data stock.Inventory
		err := rows.StructScan(&data)
		if err != nil {
			return datas, errors.Wrap(err, "[DATA][GetAllInventory]")
		}
		datas = append(datas, data)
	}
	defer rows.Close()

	return datas, nil
}

func (d Data) GetInventoryByID(ctx context.Context, id string) ([]stock.Inventory, error) {
	var (
		rows  *sqlx.Rows
		datas []stock.Inventory
		err   error
	)

	rows, err = d.stmt[getInventoryByID].QueryxContext(ctx, id)
	if err != nil {
		return datas, errors.Wrap(err, "[DATA][GetInventoryByID]")
	}

	for rows.Next() {
		var data stock.Inventory
		err := rows.StructScan(&data)
		if err != nil {
			return datas, errors.Wrap(err, "[DATA][GetInventoryByID]")
		}
		datas = append(datas, data)
	}
	defer rows.Close()

	return datas, nil
}

func (d Data) GetInventoryByIDInv(ctx context.Context, id int) (stock.Inventory, error) {
	inventory := stock.Inventory{}

	if err := d.stmt[getInventoryByIDInv].QueryRowxContext(ctx, id).StructScan(&inventory); err != nil {
		return inventory, errors.Wrap(err, "[DATA][GetInventoryByIDInv]")
	}

	return inventory, nil
}

func (d Data) CreateInventory(ctx context.Context, inventory stock.Inventory) error {
	_, err := d.stmt[createInventory].ExecContext(ctx,
		inventory.Teknisi,
		inventory.Sparepart,
		inventory.Quantity,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][CreateInventory]")
	}
	return nil
}

func (d Data) UpdateInventory(ctx context.Context, inventory stock.Inventory) error {
	_, err := d.stmt[updateInventory].ExecContext(ctx,
		inventory.Quantity,
		inventory.Teknisi,
		inventory.Sparepart,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][UpdateInventory]")
	}
	return nil
}

func (d Data) DeleteInventory(ctx context.Context, id_teknisi, id_sparepart string) error {
	_, err := d.stmt[deleteInventory].ExecContext(ctx, id_teknisi, id_sparepart)

	if err != nil {
		return errors.Wrap(err, "[DATA][DeleteInventory]")
	}
	return nil
}
