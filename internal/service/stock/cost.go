package stock

import (
	"context"
	"stock/internal/entity/stock"
)

// GetSparepartCost ...
func (s Service) GetSparepartCost(ctx context.Context, filter stock.SparepartCostFilter) ([]stock.SparepartCostResult, error) {
	// Mendapatkan hasil dari data layer
	results, err := s.data.GetSparepartCost(ctx, filter)
	if err != nil {
		return nil, err
	}

	// Memfilter hasil berdasarkan CustomerID jika ada filter
	if filter.CustomerID != nil && *filter.CustomerID != "" {
		var filteredResults []stock.SparepartCostResult
		for _, result := range results {
			if result.CustomerID == *filter.CustomerID { // Dereference pointer to get the value
				filteredResults = append(filteredResults, result)
			}
		}
		results = filteredResults
	}

	return results, nil
}
