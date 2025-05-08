package stock

import (
	"context"
	"log"
	"math"
	"stock/internal/entity/stock"
	"stock/pkg/errors"
	"time"
)

// GetAllRequests ...
func (s Service) GetAllRequests(ctx context.Context) ([]stock.Request, error) {
	request, err := s.data.GetAllRequests(ctx)
	if err != nil {
		return request, errors.Wrap(err, "[SERVICE][GetAllRequests]")
	}

	return request, nil
}

func (s Service) GetRequestsPagination(ctx context.Context, keyword, status string, page, length int) ([]stock.Request, int, error) {
	limit := length
	offset := (page - 1) * length
	var lastPage int

	if page != 0 && length != 0 {
		requests, count, err := s.data.GetRequestsCount(ctx, keyword, status)
		if err != nil {
			return requests, lastPage, errors.Wrap(err, "[SERVICE][GetRequestsFiltered][COUNT]")

		}

		lastPage = int(math.Ceil(float64(count) / float64(length)))

		requests, err = s.data.GetRequestsPage(ctx, keyword, status, offset, limit)
		if err != nil {
			return requests, lastPage, errors.Wrap(err, "[SERVICE][GetRequestsFiltered]")
		}

		return requests, lastPage, nil
	}

	requests, err := s.data.GetAllRequests(ctx)
	if err != nil {
		return requests, lastPage, errors.Wrap(err, "[SERVICE][GetRequestsFiltered]")
	}

	return requests, lastPage, nil
}

func (s Service) CreateRequest(ctx context.Context, request stock.Request) error {
	err := s.data.CreateRequest(ctx, request)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][CreateRequest]")
	}

	return nil
}

func (s Service) UpdateRequest(ctx context.Context, request stock.Request) error {
	err := s.data.UpdateRequest(ctx, request)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][UpdateRequest]")
	}

	if request.Status == "Disetujui" {
		log.Println(request)
		if request.Mesin != "" {
			history := stock.SparepartHistory{
				IDTeknisi:   request.Teknisi,
				IDMachine:   request.Mesin,
				IDSparepart: request.Sparepart,
				IDRequest:   request.ID,
				Quantity:    request.Quantity,
				// Request:     request.Counter,
				UpdatedBy:   request.UpdatedBy,
				UpdatedAt:   time.Now(),
			}
			err = s.data.CreateSparepartHistory(ctx, history)
			if err != nil {
				return errors.Wrap(err, "[SERVICE][UpdateRequest][CreateSparepartHistory]")
			}
		} else {
			inventory := stock.Inventory{
				Teknisi:   request.Teknisi,
				Sparepart: request.Sparepart,
				Quantity:  request.Quantity,
			}
			err = s.data.CreateInventory(ctx, inventory)
			if err != nil {
				return errors.Wrap(err, "[SERVICE][UpdateRequest][CreateInventory]")
			}
		}

		sparepart, err := s.data.GetSparepartByID(ctx, request.Sparepart)
		if err != nil {
			return errors.Wrap(err, "[SERVICE][UpdateRequest][GetSparepartByID]")
		}

		sparepart.Quantity = sparepart.Quantity - request.Quantity

		err = s.data.UpdateSparepart(ctx, sparepart)
		if err != nil {
			return errors.Wrap(err, "[SERVICE][UpdateRequest][UpdateSparepart]")
		}
	}

	return nil
}

func (s Service) DeleteRequest(ctx context.Context, id string) error {
	err := s.data.DeleteRequest(ctx, id)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][DeleteRequest]")
	}

	return nil
}
