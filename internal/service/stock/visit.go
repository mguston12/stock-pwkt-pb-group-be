package stock

import (
	"context"
	"log"
	"stock/internal/entity/stock"
	"stock/pkg/errors"
)

// GetAllVisits ...
func (s Service) GetVisitsByID(ctx context.Context, id int) ([]stock.Visit, error) {
	visit, err := s.data.GetVisitsByID(ctx, id)
	if err != nil {
		return visit, errors.Wrap(err, "[SERVICE][GetVisitByID]")
	}

	return visit, nil
}

func (s Service) CreateVisit(ctx context.Context, visit stock.Visit) error {
	err := s.data.CreateVisit(ctx, visit)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][CreateVisit]")
	}

	return nil
}

func (s Service) UpdateVisit(ctx context.Context, visit stock.Visit) error {
	err := s.data.UpdateVisit(ctx, visit)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][UpdateVisit]")
	}

	return nil
}

func (s Service) DeleteVisit(ctx context.Context, id int) error {
	log.Println(id)
	err := s.data.DeleteVisit(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][DeleteVisit]")
	}

	return nil
}
