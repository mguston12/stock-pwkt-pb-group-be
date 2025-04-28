package stock

import (
	"context"
	"stock/internal/entity/stock"
	"stock/pkg/errors"
)

// GetAllSparepartHistory ...
func (s Service) GetAllSparepartHistory(ctx context.Context) ([]stock.SparepartHistory, error) {
	sparepartHistory, err := s.data.GetAllSparepartHistory(ctx)
	if err != nil {
		return sparepartHistory, errors.Wrap(err, "[SERVICE][GetAllSparepartHistory]")
	}

	return sparepartHistory, nil
}
func (s Service) GetSparepartHistoryByID(ctx context.Context, id string) ([]stock.SparepartHistory, error) {
	details := []stock.SparepartHistory{}

	sparepartHistory, err := s.data.GetSparepartHistoryByID(ctx, id)
	if err != nil {
		return sparepartHistory, errors.Wrap(err, "[SERVICE][GetAllSparepartHistory]")
	}

	for _, k := range sparepartHistory {
		technician, err := s.data.GetTeknisiByID(ctx, k.IDTeknisi)
		if err != nil {
			return sparepartHistory, errors.Wrap(err, "[SERVICE][GetAllSparepartHistory]")
		}

		k.NamaTeknisi = technician.Nama

		sparepart, err := s.data.GetSparepartByID(ctx, k.IDSparepart)
		if err != nil {
			return sparepartHistory, errors.Wrap(err, "[SERVICE][GetAllSparepartHistory]")
		}

		k.NamaSparepart = sparepart.Nama

		details = append(details, k)
	}

	return details, nil
}

func (s Service) CreateSparepartHistory(ctx context.Context, history stock.SparepartHistory) error {
	err := s.data.CreateSparepartHistory(ctx, history)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][CreateSparepartHistory]")
	}

	return nil
}

func (s Service) UpdateSparepartHistory(ctx context.Context, history stock.SparepartHistory) error {
	err := s.data.UpdateSparepartHistory(ctx, history)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][UpdateSparepartHistory]")
	}

	return nil
}

func (s Service) DeleteSparepartHistory(ctx context.Context, id string) error {
	err := s.data.DeleteSparepartHistory(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][DeleteSparepartHistory]")
	}

	return nil
}
