package stock

import (
	"context"
	"math"
	"stock/internal/entity/stock"
	"stock/pkg/errors"
	"time"
)

// GetAllMachineHistories ...
func (s Service) GetAllMachineHistories(ctx context.Context, idMachine, idCustomer string, page, length int) ([]stock.MachineHistory, int, error) {
	limit := length
	offset := (page - 1) * length
	var lastPage int
	dataMachineHistory := []stock.MachineHistory{}

	if page != 0 && length != 0 {
		count, err := s.data.GetAllMachineHistoryCount(ctx, idMachine, idCustomer)
		if err != nil {
			return dataMachineHistory, lastPage, errors.Wrap(err, "[SERVICE][GetAllMachineHistories][Count]")
		}

		lastPage = int(math.Ceil(float64(count) / float64(length)))

		histories, err := s.data.GetAllMachineHistory(ctx, idMachine, idCustomer, offset, limit)
		if err != nil {
			return dataMachineHistory, lastPage, errors.Wrap(err, "[SERVICE][GetAllMachineHistories][Page]")
		}

		dataMachineHistory = histories
		return dataMachineHistory, lastPage, nil
	}

	histories, err := s.data.GetAllMachineHistory(ctx, idMachine, idCustomer, 0, 0)
	if err != nil {
		return dataMachineHistory, lastPage, errors.Wrap(err, "[SERVICE][GetAllMachineHistories][All]")
	}

	dataMachineHistory = histories
	return dataMachineHistory, lastPage, nil
}

// GetMachineHistoryByID ...
func (s Service) GetMachineHistoryByID(ctx context.Context, idMachine string) ([]stock.MachineHistory, error) {
	histories, err := s.data.GetMachineHistoryByID(ctx, idMachine)
	if err != nil {
		return histories, errors.Wrap(err, "[SERVICE][GetMachineHistoryByID]")
	}

	return histories, nil
}

// CreateMachineHistory ...
func (s Service) CreateMachineHistory(ctx context.Context, history stock.MachineHistory) error {
	err := s.data.CreateMachineHistory(ctx, history)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][CreateMachineHistory]")
	}
	return nil
}

// UpdateMachineHistory ...
func (s Service) UpdateMachineHistory(ctx context.Context, history stock.MachineHistory) error {
	err := s.data.UpdateMachineHistory(ctx, history)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][UpdateMachineHistory]")
	}
	return nil
}

// DeleteMachineHistory ...
func (s Service) DeleteMachineHistory(ctx context.Context, id int) error {
	err := s.data.DeleteMachineHistory(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][DeleteMachineHistory]")
	}
	return nil
}

func (s Service) ReplaceCustomerMachine(ctx context.Context, oldMachineID, newMachineID, customerID string) error {
	now := time.Now()

	// 1. Update history mesin lama: set tanggal_selesai & status Nonaktif
	lastHistory, err := s.data.GetLastMachineHistoryByMachineID(ctx, oldMachineID)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][ReplaceCustomerMachine][GetLastMachineHistoryByMachineID]")
	}

	lastHistory.TanggalSelesai = &now
	lastHistory.Status = "non-aktif"

	err = s.data.UpdateMachineHistory(ctx, lastHistory)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][ReplaceCustomerMachine][UpdateOldHistory]")
	}

	// 2. Update mesin baru: ubah customer jika perlu
	newMachine, err := s.data.GetMachineByID(ctx, newMachineID)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][ReplaceCustomerMachine][GetNewMachineByID]")
	}

	newMachine.Customer = customerID
	err = s.data.UpdateMachine(ctx, newMachine)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][ReplaceCustomerMachine][UpdateMachineCustomer]")
	}

	// 3. Buat entry baru di history untuk mesin baru
	newHistory := stock.MachineHistory{
		IDMachine:      newMachineID,
		TanggalMulai:   now,
		TanggalSelesai: nil,
		Status:         "aktif",
	}

	err = s.data.CreateMachineHistory(ctx, newHistory)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][ReplaceCustomerMachine][CreateNewMachineHistory]")
	}

	return nil
}

// DeactivateMachine digunakan untuk menonaktifkan mesin
func (s Service) DeactivateMachine(ctx context.Context, machineID string) error {
	// Ambil history aktif terakhir untuk mesin
	history, err := s.data.GetLastMachineHistoryByMachineID(ctx, machineID)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][DeactivateMachine][GetLastMachineHistoryByMachineID]")
	}

	// Jika sudah non-aktif, tidak perlu diproses
	if history.Status == "non-aktif" {
		return errors.New("mesin sudah berstatus non-aktif")
	}

	// Set tanggal_selesai dan status
	now := time.Now()
	history.TanggalSelesai = &now
	history.Status = "non-aktif"

	// Update ke database
	err = s.data.UpdateMachineHistory(ctx, history)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][DeactivateMachine][UpdateMachineHistory]")
	}

	return nil
}
