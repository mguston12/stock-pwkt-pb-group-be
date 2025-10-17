package stock

import (
	"context"
	"database/sql"
	"stock-pwt/internal/entity/stock"

	"stock-pwt/pkg/errors"
)

// Transaction
func (d Data) GetInventoryByIDInvTx(ctx context.Context, tx *sql.Tx, id int) (stock.Inventory, error) {
	query := `SELECT i.id_inventory, i.id_teknisi, i.id_sparepart, i.quantity, sp.nama_sparepart 
			  FROM inventory i
			  LEFT JOIN sparepart sp ON i.id_sparepart = sp.id_sparepart 
			  WHERE i.id_inventory = ?`

	row := tx.QueryRowContext(ctx, query, id)

	var inv stock.Inventory
	err := row.Scan(&inv.ID, &inv.Teknisi, &inv.Sparepart, &inv.Quantity, &inv.NamaSparepart)
	if err != nil {
		return inv, errors.Wrap(err, "[DATA][GetInventoryByIDInvTx]")
	}

	return inv, nil
}

func (d Data) UpdateInventoryTx(ctx context.Context, tx *sql.Tx, inv stock.Inventory) error {
	query := `
		UPDATE inventory 
		SET quantity = ?
		WHERE id_teknisi = ? AND id_sparepart = ?
	`

	_, err := tx.ExecContext(ctx, query, inv.Quantity, inv.Teknisi, inv.Sparepart)
	if err != nil {
		return errors.Wrap(err, "[DATA][UpdateInventoryTx]")
	}

	return nil
}

func (d Data) CreateSparepartHistoryTx(ctx context.Context, tx *sql.Tx, history stock.SparepartHistory) error {
	query := `
		INSERT INTO history_sparepart 
		(id_teknisi, id_machine, id_sparepart, quantity, counter, counter_colour, counter_colour_a3, updated_by)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`

	_, err := tx.ExecContext(ctx, query,
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
		return errors.Wrap(err, "[DATA][CreateSparepartHistoryTx]")
	}

	return nil
}
