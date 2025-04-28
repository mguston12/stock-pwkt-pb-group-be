package stock

import (
	"context"
	"stock/internal/entity/stock"
)

// Data ...
// Masukkan function dari package data ke dalam interface ini
type Data interface {
	GetAllSpareparts(ctx context.Context) ([]stock.Sparepart, error)
	GetAllSparepartsPage(ctx context.Context, keyword string, offset, limit int) ([]stock.Sparepart, error)
	GetAllSparepartsCount(ctx context.Context, keyword string) ([]stock.Sparepart, int, error)
	GetSparepartByID(ctx context.Context, id string) (stock.Sparepart, error)
	CreateSparepart(ctx context.Context, sparepart stock.Sparepart) error
	UpdateSparepart(ctx context.Context, sparepart stock.Sparepart) error
	DeleteSparepart(ctx context.Context, id string) error

	GetAllSparepartHistory(ctx context.Context) ([]stock.SparepartHistory, error)
	GetSparepartHistoryByID(ctx context.Context, id string) ([]stock.SparepartHistory, error)
	CreateSparepartHistory(ctx context.Context, history stock.SparepartHistory) error
	UpdateSparepartHistory(ctx context.Context, history stock.SparepartHistory) error
	DeleteSparepartHistory(ctx context.Context, id string) error

	GetAllTeknisi(ctx context.Context) ([]stock.Teknisi, error)
	GetTeknisiByID(ctx context.Context, id string) (stock.Teknisi, error)
	CreateTeknisi(ctx context.Context, teknisi stock.Teknisi) error
	UpdateTeknisi(ctx context.Context, teknisi stock.Teknisi) error
	DeleteTeknisi(ctx context.Context, id string) error

	GetAllMachines(ctx context.Context) ([]stock.Machine, error)
	GetMachineByID(ctx context.Context, id string) (stock.Machine, error)
	CreateMachine(ctx context.Context, machine stock.Machine) error
	UpdateMachine(ctx context.Context, machine stock.Machine) error
	DeleteMachine(ctx context.Context, id string) error

	GetAllCustomers(ctx context.Context) ([]stock.Customer, error)
	CreateCustomer(ctx context.Context, customer stock.Customer) error
	UpdateCustomer(ctx context.Context, customer stock.Customer) error
	DeleteCustomer(ctx context.Context, id string) error

	GetAllRequests(ctx context.Context) ([]stock.Request, error)
	GetRequestsPage(ctx context.Context, offset, limit int) ([]stock.Request, error)
	GetRequestsCount(ctx context.Context) ([]stock.Request, int, error)
	CreateRequest(ctx context.Context, request stock.Request) error
	UpdateRequest(ctx context.Context, request stock.Request) error
	DeleteRequest(ctx context.Context, id string) error

	GetUserByUsername(ctx context.Context, username string) (stock.User, error)
	CreateUser(ctx context.Context, user stock.User) error
	UpdateUser(ctx context.Context, user stock.User) error
	DeleteUser(ctx context.Context, username string) error
}

// Service ...
// Tambahkan variable sesuai banyak data layer yang dibutuhkan
type Service struct {
	data Data
}

// New ...
// Tambahkan parameter sesuai banyak data layer yang dibutuhkan
func New(data Data) Service {
	// Assign variable dari parameter ke object
	return Service{
		data: data,
	}
}
