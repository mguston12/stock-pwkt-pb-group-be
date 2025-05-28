package stock

import (
	"context"
	"math"
	"stock/internal/entity/stock"
	"stock/pkg/errors"
)

// GetAllMachines ...
func (s Service) GetAllMachines(ctx context.Context) ([]stock.Machine, error) {
	machine, err := s.data.GetAllMachines(ctx)
	if err != nil {
		return machine, errors.Wrap(err, "[SERVICE][GetAllMachines]")
	}

	return machine, nil
}

func (s Service) GetMachinesFiltered(ctx context.Context, keyword string, page, length int) ([]stock.Machine, int, error) {
	limit := length
	offset := (page - 1) * length
	var lastPage int

	if page != 0 && length != 0 {
		machines, count, err := s.data.GetAllMachinesCount(ctx, keyword)
		if err != nil {
			return machines, lastPage, errors.Wrap(err, "[SERVICE][GetMachinesFiltered][COUNT]")

		}

		lastPage = int(math.Ceil(float64(count) / float64(length)))

		machines, err = s.data.GetAllMachinesPage(ctx, keyword, offset, limit)
		if err != nil {
			return machines, lastPage, errors.Wrap(err, "[SERVICE][GetMachinesFiltered]")
		}

		return machines, lastPage, errors.Wrap(err, "[SERVICE][GetMachinesFiltered]")
	}

	machines, err := s.data.GetAllMachines(ctx)
	if err != nil {
		return machines, lastPage, errors.Wrap(err, "[SERVICE][GetMachinesFiltered]")
	}

	return machines, lastPage, nil
}

func (s Service) GetMachineByID(ctx context.Context, id string) (stock.Machine, error) {
	machine, err := s.data.GetMachineByID(ctx, id)
	if err != nil {
		return machine, errors.Wrap(err, "[SERVICE][GetMachineByID]")
	}

	history, err := s.GetSparepartHistoryByID(ctx, id)
	if err != nil {
		return machine, errors.Wrap(err, "[SERVICE][GetMachineByID]")
	}

	if len(history) > 0 {
		machine.History = history
	}

	return machine, nil
}

func (s Service) GetMachineByIDCustomer(ctx context.Context, id string) ([]stock.Machine, error) {
	machine, err := s.data.GetMachineByIDCustomer(ctx, id)
	if err != nil {
		return machine, errors.Wrap(err, "[SERVICE][GetMachineByIDCustomer]")
	}

	return machine, nil
}

func (s Service) CreateMachine(ctx context.Context, machine stock.Machine) error {
	err := s.data.CreateMachine(ctx, machine)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][CreateMachine]")
	}

	return nil
}

func (s Service) UpdateMachine(ctx context.Context, machine stock.Machine) error {
	err := s.data.UpdateMachine(ctx, machine)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][UpdateMachine]")
	}

	return nil
}

func (s Service) DeleteMachine(ctx context.Context, id string) error {
	err := s.data.DeleteMachine(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][DeleteMachine]")
	}

	return nil
}
