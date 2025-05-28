package stock

import (
	"context"
	"fmt"
	"stock/internal/entity/stock"
)

func (d Data) GetSparepartHistoryByMonth(ctx context.Context, bulan, tahun int) ([]stock.ReportData, error) {
	query := `SELECT hs.id_history, hs.id_teknisi, hs.id_machine, hs.id_sparepart, sp.nama_sparepart, hs.quantity,
				hs.counter, hs.updated_by, hs.updated_at, t.nama_teknisi, c.nama_customer, 
				COALESCE(avg_cost.average_cost, 0) AS average_cost
				FROM history_sparepart hs
				JOIN 
					machine_history mh ON hs.id_machine = mh.id_machine
				JOIN 
					sparepart sp ON hs.id_sparepart = sp.id_sparepart
				JOIN 
					teknisi t ON hs.id_teknisi = t.id_teknisi
				JOIN 
					customer c ON mh.id_customer = c.id_customer
				LEFT JOIN (
				SELECT 
					id_sparepart, 
					COALESCE(SUM(quantity * harga_per_unit) / NULLIF(SUM(quantity), 0), 0) AS average_cost
				FROM 
					pembelian_sparepart
				GROUP BY 
					id_sparepart
				) AS avg_cost ON hs.id_sparepart = avg_cost.id_sparepart
				WHERE 
					hs.updated_at BETWEEN mh.tanggal_mulai AND COALESCE(mh.tanggal_selesai, NOW())
				AND 
					MONTH(hs.updated_at) = ?  AND YEAR(hs.updated_at) = ?
				ORDER BY t.nama_teknisi, sp.nama_sparepart ASC`

	rows, err := d.db.QueryxContext(ctx, query, bulan, tahun)
	if err != nil {
		return nil, fmt.Errorf("[DATA][GetSparepartHistoryByMonth] gagal query: %w", err)
	}
	defer rows.Close()

	var results []stock.ReportData
	for rows.Next() {
		var h stock.ReportData
		if err := rows.StructScan(&h); err != nil {
			return nil, fmt.Errorf("[DATA][GetSparepartHistoryByMonth] gagal scan: %w", err)
		}
		results = append(results, h)
	}

	return results, nil
}
