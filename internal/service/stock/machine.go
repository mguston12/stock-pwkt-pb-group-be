package stock

import (
	"context"
	"fmt"
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

func (s Service) GetMachineByID(ctx context.Context, id string) (stock.Machine, error) {
	machine, err := s.data.GetMachineByID(ctx, id)
	if err != nil {
		return machine, errors.Wrap(err, "[SERVICE][GetMachineByID]")
	}

	history, err := s.GetSparepartHistoryByID(ctx, id)
	if err != nil {
		return machine, errors.Wrap(err, "[SERVICE][GetMachineByID]")
	}

	fmt.Println(history)

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
