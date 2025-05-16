package stock

import (
	"context"
	"stock/internal/entity/stock"
	"stock/pkg/errors"
)

func (d Data) GetSparepartCost(ctx context.Context, filter stock.SparepartCostFilter) ([]stock.SparepartCostResult, error) {
	query := `
	WITH average_costs AS (
    SELECT 
        id_sparepart, 
        SUM(quantity * harga_per_unit) / SUM(quantity) AS average_cost
    FROM 
        pembelian_sparepart
    GROUP BY 
        id_sparepart
	)
	SELECT 
		c.id_customer, 
		c.nama_customer,
		m.id_machine,
		DATE_FORMAT(h.updated_at, '%Y-%m') AS bulan,
		YEAR(h.updated_at) AS tahun,
		SUM(h.quantity * ac.average_cost) AS total_cost
	FROM 
		history_sparepart h
	JOIN 
		machine m ON h.id_machine = m.id_machine
	JOIN 
		customer c ON m.id_customer = c.id_customer
	JOIN 
		average_costs ac ON h.id_sparepart = ac.id_sparepart
	WHERE 
		(? IS NULL OR YEAR(h.updated_at) = ?) AND
		(? IS NULL OR MONTH(h.updated_at) = ?) AND
		(? IS NULL OR c.id_customer = ?) AND
		(? IS NULL OR m.id_machine = ?)
	GROUP BY 
		c.id_customer, c.nama_customer, m.id_machine, YEAR(h.updated_at), MONTH(h.updated_at), h.updated_at
	ORDER BY 
		tahun, bulan, c.nama_customer, m.id_machine`

	rows, err := d.db.QueryxContext(ctx, query,
		filter.Year, filter.Year,
		filter.Month, filter.Month,
		filter.CustomerID, filter.CustomerID,
		filter.MachineID, filter.MachineID,
	)
	if err != nil {
		return nil, errors.Wrap(err, "[DATA][GetSparepartCost]")
	}
	defer rows.Close()

	var results []stock.SparepartCostResult
	for rows.Next() {
		var r stock.SparepartCostResult
		if err := rows.StructScan(&r); err != nil {
			return nil, errors.Wrap(err, "[DATA][GetSparepartCost] scan result")
		}
		results = append(results, r)
	}
	return results, nil
}
