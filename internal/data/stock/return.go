package stock

import (
	"context"
	"stock/internal/entity/stock"
	"stock/pkg/errors"
)

// Get all return inventory
func (d Data) GetAllReturnInventory(ctx context.Context) ([]stock.ReturnInventory, error) {
	var result []stock.ReturnInventory

	rows, err := d.stmt[getAllReturnInventory].QueryxContext(ctx)
	if err != nil {
		return result, errors.Wrap(err, "[DATA][GetAllReturnInventory]")
	}
	defer rows.Close()

	for rows.Next() {
		var data stock.ReturnInventory
		if err := rows.StructScan(&data); err != nil {
			return result, errors.Wrap(err, "[DATA][GetAllReturnInventory][StructScan]")
		}
		result = append(result, data)
	}

	return result, nil
}

// Get return inventory by status
func (d Data) GetReturnInventoryByStatus(ctx context.Context, status string) ([]stock.ReturnInventory, error) {
	var result []stock.ReturnInventory

	rows, err := d.stmt[getReturnInventoryByStatus].QueryxContext(ctx, status)
	if err != nil {
		return result, errors.Wrap(err, "[DATA][GetReturnInventoryByStatus]")
	}
	defer rows.Close()

	for rows.Next() {
		var data stock.ReturnInventory
		if err := rows.StructScan(&data); err != nil {
			return result, errors.Wrap(err, "[DATA][GetReturnInventoryByStatus][StructScan]")
		}
		result = append(result, data)
	}

	return result, nil
}

// Create return inventory record
func (d Data) CreateReturnInventory(ctx context.Context, input stock.ReturnInventory) error {
	_, err := d.stmt[createReturnInventory].ExecContext(ctx,
		input.IDInventory,
		input.IDSparepart,
		input.Quantity,
		input.ReturnedBy,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][CreateReturnInventory]")
	}

	return nil
}

// Approve or reject return
func (d Data) ApproveReturnInventory(ctx context.Context, input stock.ReturnInventory) error {
	_, err := d.stmt[approveReturnInventory].ExecContext(ctx,
		input.Status,
		input.ApprovedBy,
		input.IDReturn,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][ApproveReturnInventory]")
	}

	return nil
}
