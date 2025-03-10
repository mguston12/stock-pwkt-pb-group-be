package stock

import (
	"context"
	"stock/internal/entity/stock"
	"stock/pkg/errors"
)

// GetAllSpareparts ...
func (s Service) GetAllSpareparts(ctx context.Context) ([]stock.Sparepart, error) {
	sparepart, err := s.data.GetAllSpareparts(ctx)
	if err != nil {
		return sparepart, errors.Wrap(err, "[SERVICE][GetAllSpareparts]")
	}

	return sparepart, nil
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
	err := s.data.DeleteSparepart(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][DeleteSparepart]")
	}

	return nil
}
