package stock

import (
	"context"
	"log"
	"stock/internal/entity/stock"
	"stock/pkg/errors"
)

func (s Service) GetPembelianSparepart(ctx context.Context) ([]stock.PembelianSparepart, error) {
	pembelian_sp, err := s.data.GetPembelianSparepart(ctx)
	if err != nil {
		return pembelian_sp, errors.Wrap(err, "[SERVICE][GetPembelianSparepart]")
	}

	return pembelian_sp, nil
}

func (s Service) GetPembelianSparepartByID(ctx context.Context, id string) ([]stock.PembelianSparepart, error) {
	pembelian_sp, err := s.data.GetPembelianSparepartByID(ctx, id)
	if err != nil {
		return pembelian_sp, errors.Wrap(err, "[SERVICE][GetPembelianSparepartByID]")
	}

	return pembelian_sp, nil
}

func (s Service) GetPembelianSparepartBySupplier(ctx context.Context, id string) ([]stock.PembelianSparepart, error) {
	pembelian_sp, err := s.data.GetPembelianSparepartBySupplier(ctx, id)
	if err != nil {
		return pembelian_sp, errors.Wrap(err, "[SERVICE][GetPembelianSparepartBySupplier]")
	}

	return pembelian_sp, nil
}

func (s Service) CreatePembelianSparepart(ctx context.Context, pembelian_sp stock.PembelianSparepart) error {
	err := s.data.CreatePembelianSparepart(ctx, pembelian_sp)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][CreatePembelianSparepart]")
	}

	sparepart, err := s.data.GetSparepartByID(ctx, pembelian_sp.Sparepart)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][CreatePembelianSparepart][2]")
	}

	sp := stock.Sparepart{
		Nama:     sparepart.Nama,
		ID:       sparepart.ID,
		Quantity: sparepart.Quantity + pembelian_sp.Quantity,
	}

	err = s.data.UpdateSparepart(ctx, sp)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][CreatePembelianSparepart][3]")
	}

	return nil
}

func (s Service) UpdatePembelianSparepart(ctx context.Context, pembelian_sp stock.PembelianSparepart) error {
	err := s.data.UpdatePembelianSparepart(ctx, pembelian_sp)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][UpdatePembelianSparepart]")
	}

	return nil
}

func (s Service) DeletePembelianSparepart(ctx context.Context, id string) error {
	err := s.data.DeletePembelianSparepart(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][DeletePembelianSparepart]")
	}

	return nil
}
