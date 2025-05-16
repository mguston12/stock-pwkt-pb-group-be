package stock

import (
	"context"
	"log"
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

func (s Service) InventoryUsage(ctx context.Context, input stock.InventoryUsage) error {
	inventory, err := s.data.GetInventoryByIDInv(ctx, input.InventoryID)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][InventoryUsage][1]")
	}

	inv := stock.Inventory{
		ID:            input.InventoryID,
		Teknisi:       input.UpdatedBy,
		Sparepart:     inventory.Sparepart,
		NamaSparepart: inventory.NamaSparepart,
		Quantity:      inventory.Quantity - 1,
	}

	log.Println(inv)

	err = s.data.UpdateInventory(ctx, inv)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][InventoryUsage][2]")
	}

	history := stock.SparepartHistory{
		IDTeknisi:   input.UpdatedBy,
		IDMachine:   input.MachineID,
		IDSparepart: inventory.Sparepart,
		Quantity:    1,
		Counter:     input.Counter,
		UpdatedBy:   input.UpdatedBy,
	}

	err = s.data.CreateSparepartHistory(ctx, history)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][InventoryUsage][3]")
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
