package stock

import (
	"context"
	"math"
	"stock-pwt/internal/entity/stock"
	"stock-pwt/pkg/errors"
)

// GetAllSparepartHistory ...
// func (s Service) GetAllSparepartHistory(ctx context.Context) ([]stock.SparepartHistory, error) {
// 	sparepartHistory, err := s.data.GetAllSparepartHistory(ctx)
// 	if err != nil {
// 		return sparepartHistory, errors.Wrap(err, "[SERVICE][GetAllSparepartHistory]")
// 	}

// 	return sparepartHistory, nil
// }

func (s Service) GetAllSparepartHistory(ctx context.Context, id_teknisi, id_mesin, id_sparepart string, page, length int) ([]stock.SparepartHistory, int, error) {
	limit := length
	offset := (page - 1) * length
	var lastPage int
	details := []stock.SparepartHistory{}

	histories, count, err := s.data.GetAllSparepartHistoryCount(ctx, id_teknisi, id_mesin, id_sparepart)
	if err != nil {
		return histories, lastPage, errors.Wrap(err, "[SERVICE][GetSparepartsFiltered][COUNT]")

	}

	lastPage = int(math.Ceil(float64(count) / float64(length)))

	histories, err = s.data.GetAllSparepartHistory(ctx, id_teknisi, id_mesin, id_sparepart, offset, limit)
	if err != nil {
		return histories, lastPage, errors.Wrap(err, "[SERVICE][GetSparepartsFiltered]")
	}

	for _, k := range histories {
		technician, err := s.data.GetTeknisiByID(ctx, k.IDTeknisi)
		if err != nil {
			return histories, lastPage, errors.Wrap(err, "[SERVICE][GetAllSparepartHistory]")
		}

		k.NamaTeknisi = technician.Nama

		sparepart, err := s.data.GetSparepartByID(ctx, k.IDSparepart)
		if err != nil {
			return histories, lastPage, errors.Wrap(err, "[SERVICE][GetAllSparepartHistory]")
		}

		k.NamaSparepart = sparepart.Nama

		details = append(details, k)
	}

	return details, lastPage, nil
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
