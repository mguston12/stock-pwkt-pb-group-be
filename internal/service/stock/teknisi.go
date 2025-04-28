package stock

import (
	"context"
	"stock/internal/entity/stock"
	"stock/pkg/errors"
)

// GetAllTeknisis ...
func (s Service) GetAllTeknisi(ctx context.Context) ([]stock.Teknisi, error) {
	Teknisi, err := s.data.GetAllTeknisi(ctx)
	if err != nil {
		return Teknisi, errors.Wrap(err, "[SERVICE][GetAllTeknisi]")
	}

	return Teknisi, nil
}

func (s Service) GetTeknisiByID(ctx context.Context, id string) (stock.Teknisi, error) {
	teknisi, err := s.data.GetTeknisiByID(ctx, id)
	if err != nil {
		return teknisi, errors.Wrap(err, "[SERVICE][GetTeknisiByID]")
	}

	return teknisi, nil
}

func (s Service) CreateTeknisi(ctx context.Context, teknisi stock.Teknisi) error {
	err := s.data.CreateTeknisi(ctx, teknisi)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][CreateTeknisi]")
	}

	return nil
}

func (s Service) UpdateTeknisi(ctx context.Context, teknisi stock.Teknisi) error {
	err := s.data.UpdateTeknisi(ctx, teknisi)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][UpdateTeknisi]")
	}

	return nil
}

func (s Service) DeleteTeknisi(ctx context.Context, id string) error {
	err := s.data.DeleteTeknisi(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][DeleteTeknisi]")
	}

	return nil
}
