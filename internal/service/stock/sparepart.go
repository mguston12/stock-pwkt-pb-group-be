package stock

import (
	"context"
	"math"
	"stock-pwt/internal/entity/stock"
	"stock-pwt/pkg/errors"
)

// GetAllSpareparts ...
func (s Service) GetAllSpareparts(ctx context.Context) ([]stock.Sparepart, error) {
	sparepart, err := s.data.GetAllSpareparts(ctx)
	if err != nil {
		return sparepart, errors.Wrap(err, "[SERVICE][GetAllSpareparts]")
	}

	return sparepart, nil
}

func (s Service) GetSparepartsFiltered(ctx context.Context, keyword string, page, length int) ([]stock.Sparepart, int, error) {
	limit := length
	offset := (page - 1) * length
	var lastPage int
	dataSparepart := []stock.Sparepart{}

	if page != 0 && length != 0 {
		spareparts, count, err := s.data.GetAllSparepartsCount(ctx, keyword)
		if err != nil {
			return dataSparepart, lastPage, errors.Wrap(err, "[SERVICE][GetSparepartsFiltered][COUNT]")

		}

		lastPage = int(math.Ceil(float64(count) / float64(length)))

		spareparts, err = s.data.GetAllSparepartsPage(ctx, keyword, offset, limit)
		if err != nil {
			return dataSparepart, lastPage, errors.Wrap(err, "[SERVICE][GetSparepartsFiltered]")
		}

		for _, sp := range spareparts {
			cost, err := s.data.GetAverageCostSparepart(ctx, sp.ID)
			if err != nil {
				return dataSparepart, lastPage, errors.Wrap(err, "[SERVICE][GetSparepartsFiltered]")
			}

			sp = stock.Sparepart{
				ID:          sp.ID,
				Nama:        sp.Nama,
				Quantity:    sp.Quantity,
				AverageCost: cost,
			}

			dataSparepart = append(dataSparepart, sp)
		}

		return dataSparepart, lastPage, nil
	}

	spareparts, err := s.data.GetAllSpareparts(ctx)
	if err != nil {
		return spareparts, lastPage, errors.Wrap(err, "[SERVICE][GetSparepartsFilteredf]")
	}

	for _, sp := range spareparts {
		cost, err := s.data.GetAverageCostSparepart(ctx, sp.ID)
		if err != nil {
			return dataSparepart, lastPage, errors.Wrap(err, "[SERVICE][GetSparepartsFiltered]")
		}

		sp = stock.Sparepart{
			ID:          sp.ID,
			Nama:        sp.Nama,
			Quantity:    sp.Quantity,
			AverageCost: cost,
		}

		dataSparepart = append(dataSparepart, sp)
	}
	return dataSparepart, lastPage, nil
}

func (s Service) CreateSparepart(ctx context.Context, sparepart stock.Sparepart) error {
	err := s.data.CreateSparepart(ctx, sparepart)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][CreateSparepart]")
	}

	return nil
}

func (s Service) UpdateSparepart(ctx context.Context, sparepart stock.Sparepart) error {
	err := s.data.UpdateSparepart(ctx, sparepart)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][UpdateSparepart]")
	}

	return nil
}

func (s Service) DeleteSparepart(ctx context.Context, id string) error {
	count, err := s.data.CheckSparepartValidOrNot(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][DeleteSparepart][1]")
	}

	if count == 0 {
		err = s.data.DeletePembelianSparepartByIDSparepart(ctx, id)
		if err != nil {
			return errors.Wrap(err, "[SERVICE][DeleteSparepart][2]")
		}

		err = s.data.DeleteRequestByIDSparepart(ctx, id)
		if err != nil {
			return errors.Wrap(err, "[SERVICE][DeleteSparepart][3]")
		}

		err = s.data.DeleteSparepart(ctx, id)
		if err != nil {
			return errors.Wrap(err, "[SERVICE][DeleteSparepart][4]")
		}

	} else {
		return errors.New("Sparepart tidak bisa dihapus karena sudah digunakan dalam request yang disetujui")

	}

	return nil
}
