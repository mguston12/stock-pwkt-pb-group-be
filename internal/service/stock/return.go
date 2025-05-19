package stock

import (
	"context"
	"log"
	"stock/internal/entity/stock"
	"stock/pkg/errors"
)

// GetAllReturnInventory ...
func (s Service) GetAllReturnInventory(ctx context.Context) ([]stock.ReturnInventory, error) {
	returnInventories, err := s.data.GetAllReturnInventory(ctx)
	if err != nil {
		return returnInventories, errors.Wrap(err, "[SERVICE][GetAllReturnInventory]")
	}
	return returnInventories, nil
}

// GetReturnInventoryByStatus ...
func (s Service) GetReturnInventoryByStatus(ctx context.Context, status string) ([]stock.ReturnInventory, error) {
	returnInventories, err := s.data.GetReturnInventoryByStatus(ctx, status)
	if err != nil {
		return returnInventories, errors.Wrap(err, "[SERVICE][GetReturnInventoryByStatus]")
	}
	return returnInventories, nil
}

// ProcessReturnSparepart - Fungsi ini mencatat pengembalian sparepart dan memperbarui inventory
func (s Service) ProcessReturnSparepart(ctx context.Context, input stock.ReturnInventory) error {
	err := s.data.CreateReturnInventory(ctx, input)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][CreateReturnInventory]")
	}

	inv, err := s.data.GetInventoryByIDInv(ctx, input.IDInventory)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][GetInventoryByIDInv]")
	}

	if inv.Quantity < input.Quantity {
		return errors.New("[SERVICE][ProcessReturnSparepart] Quantity returned exceeds available stock")
	}

	updatedInventory := stock.Inventory{
		ID:        input.IDInventory,
		Sparepart: inv.Sparepart,
		Quantity:  inv.Quantity - input.Quantity,
		Teknisi:   input.ReturnedBy,
	}

	log.Println(updatedInventory)

	err = s.data.UpdateInventory(ctx, updatedInventory)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][UpdateInventory]")
	}

	return nil
}

// ApproveReturnInventory ...
func (s Service) ApproveReturnInventory(ctx context.Context, input stock.ReturnInventory) error {
	// 1. Update status return inventory ke database (untuk mencatat bahwa pengembalian disetujui atau ditolak)
	err := s.data.ApproveReturnInventory(ctx, input)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][ApproveReturnInventory]")
	}

	// 2. Jika status pengembalian adalah "Disetujui", kita akan menambah stok sparepart
	if input.Status == "Disetujui" {
		// Ambil data sparepart berdasarkan IDSparepart
		sparepart, err := s.data.GetSparepartByID(ctx, input.IDSparepart)
		if err != nil {
			return errors.Wrap(err, "[SERVICE][GetSparepartByID] during approval")
		}

		// Update stok sparepart dengan menambah quantity
		updatedSparepart := stock.Sparepart{
			ID:       sparepart.ID,
			Nama:     sparepart.Nama,
			Quantity: sparepart.Quantity + input.Quantity, // Menambah quantity sparepart
		}

		// Update database untuk sparepart
		err = s.data.UpdateSparepart(ctx, updatedSparepart)
		if err != nil {
			return errors.Wrap(err, "[SERVICE][UpdateSparepart] during approval")
		}

	} else if input.Status == "Ditolak" {
		// Jika statusnya Ditolak, kita akan mengembalikan sparepart ke inventory teknisi
		// Ambil data inventory berdasarkan IDInventory
		inv, err := s.data.GetInventoryByIDInv(ctx, input.IDInventory)
		if err != nil {
			return errors.Wrap(err, "[SERVICE][GetInventoryByIDInv] during rejection")
		}

		// Update inventory teknisi dengan menambah quantity sparepart yang ditolak
		updatedInventory := stock.Inventory{
			ID:        inv.ID,
			Teknisi:   inv.Teknisi,
			Sparepart: inv.Sparepart,
			Quantity:  inv.Quantity + input.Quantity, // Menambah quantity inventory karena retur ditolak
		}

		// Update inventory
		err = s.data.UpdateInventory(ctx, updatedInventory)
		if err != nil {
			return errors.Wrap(err, "[SERVICE][UpdateInventory] during rejection")
		}
	}

	return nil
}
