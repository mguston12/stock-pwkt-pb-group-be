package stock

import (
	"context"
	"log"
	"stock/internal/entity/stock"

	"stock/pkg/errors"

	"github.com/jmoiron/sqlx"
)

func (d Data) GetPembelianSparepart(ctx context.Context) ([]stock.PembelianSparepart, error) {
	var (
		rows  *sqlx.Rows
		datas []stock.PembelianSparepart
		err   error
	)

	rows, err = d.stmt[getPembelianSparepart].QueryxContext(ctx)
	if err != nil {
		return datas, errors.Wrap(err, "[DATA][GetPembelianSparepart]")
	}

	for rows.Next() {
		var data stock.PembelianSparepart
		err := rows.StructScan(&data)
		if err != nil {
			return datas, errors.Wrap(err, "[DATA][GetPembelianSparepart]")
		}
		datas = append(datas, data)
	}
	defer rows.Close()

	return datas, nil
}

func (d Data) GetPembelianSparepartByID(ctx context.Context, id string) ([]stock.PembelianSparepart, error) {
	var (
		rows  *sqlx.Rows
		datas []stock.PembelianSparepart
		err   error
	)

	rows, err = d.stmt[getPembelianSparepartByID].QueryxContext(ctx, id)
	if err != nil {
		return datas, errors.Wrap(err, "[DATA][GetPembelianSparepartByID]")
	}

	for rows.Next() {
		var data stock.PembelianSparepart
		err := rows.StructScan(&data)
		if err != nil {
			return datas, errors.Wrap(err, "[DATA][GetPembelianSparepartByID]")
		}
		datas = append(datas, data)
	}
	defer rows.Close()

	return datas, nil
}

func (d Data) GetAverageCostSparepart(ctx context.Context, id string) (float64, error) {
	averagecost := 0.0

	if err := d.stmt[getAverageCostSparepart].QueryRowxContext(ctx, id).Scan(&averagecost); err != nil {
		return averagecost, errors.Wrap(err, "[DATA][GetAverageCostSparepart]")
	}

	return averagecost, nil
}

func (d Data) CreatePembelianSparepart(ctx context.Context, pembelian_sp stock.PembelianSparepart) error {
	_, err := d.stmt[createPembelianSparepart].ExecContext(ctx,
		pembelian_sp.Sparepart,
		pembelian_sp.Quantity,
		pembelian_sp.HargaPerUnit,
		pembelian_sp.Supplier,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][CreatePembelianSparepart]")
	}
	return nil
}

func (d Data) UpdatePembelianSparepart(ctx context.Context, pembelian_sp stock.PembelianSparepart) error {
	_, err := d.stmt[updatePembelianSparepart].ExecContext(ctx,
		pembelian_sp.Quantity,
		pembelian_sp.HargaPerUnit,
		pembelian_sp.ID,
	)

	if err != nil {
		return errors.Wrap(err, "[DATA][UpdatePembelianSparepart]")
	}
	return nil
}

func (d Data) DeletePembelianSparepart(ctx context.Context, id string) error {
	result, err := d.stmt[deletePembelianSparepart].ExecContext(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[DATA][DeletePembelianSparepart]")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "[DATA][DeletePembelianSparepart][RowsAffected]")
	}

	if rowsAffected == 0 {
		log.Printf("[INFO] Tidak ada pembelian sparepart dengan id %s untuk dihapus", id)
	}

	return nil
}

func (d Data) DeletePembelianSparepartByIDSparepart(ctx context.Context, id string) error {
	result, err := d.stmt[deletePembelianSparepartByIDSparepart].ExecContext(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[DATA][DeletePembelianSparepartByIDSparepart]")
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return errors.Wrap(err, "[DATA][DeletePembelianSparepartByIDSparepart][RowsAffected]")
	}

	if rowsAffected == 0 {
		log.Printf("[INFO] Tidak ada pembelian sparepart dengan id %s untuk dihapus", id)
	}

	return nil
}
