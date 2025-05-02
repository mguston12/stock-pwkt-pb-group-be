package stock

import (
	"context"
	"stock/internal/entity/stock"
	"stock/pkg/errors"
)

func (s Service) GetAllInventory(ctx context.Context) ([]stock.Inventory, error) {
	inventory, err := s.data.GetAllInventory(ctx)
	if err != nil {
		return inventory, errors.Wrap(err, "[SERVICE][GetAllInventory]")
	}

	return inventory, nil
}

func (s Service) GetInventoryByID(ctx context.Context, id string) ([]stock.Inventory, error) {
	inventory, err := s.data.GetInventoryByID(ctx, id)
	if err != nil {
		return inventory, errors.Wrap(err, "[SERVICE][GetInventoryByID]")
	}

	return inventory, nil
}

func (s Service) CreateInventory(ctx context.Context, inventory stock.Inventory) error {
	err := s.data.CreateInventory(ctx, inventory)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][CreateInventory]")
	}

	return nil
}

func (s Service) UpdateInventory(ctx context.Context, inventory stock.Inventory) error {
	err := s.data.UpdateInventory(ctx, inventory)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][UpdateInventory]")
	}

	return nil
}

func (s Service) DeleteInventory(ctx context.Context, id_teknisi, id_sparepart string) error {
	err := s.data.DeleteInventory(ctx, id_teknisi, id_sparepart)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][DeleteInventory]")
	}

	return nil
}
