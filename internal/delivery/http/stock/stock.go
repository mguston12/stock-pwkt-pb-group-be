package stock

import (
	"context"
	"stock/internal/entity/stock"
)

type StockSvc interface {
	GetAllSpareparts(ctx context.Context) ([]stock.Sparepart, error)
	GetSparepartsFiltered(ctx context.Context, keyword string, page, length int) ([]stock.Sparepart, int, error)
	CreateSparepart(ctx context.Context, sparepart stock.Sparepart) error
	UpdateSparepart(ctx context.Context, sparepart stock.Sparepart) error
	DeleteSparepart(ctx context.Context, id string) error
	GetSparepartCost(ctx context.Context, filter stock.SparepartCostFilter) ([]stock.SparepartCostResult, error)

	GetAllTeknisi(ctx context.Context) ([]stock.Teknisi, error)
	GetTeknisiByID(ctx context.Context, id string) (stock.Teknisi, error)
	CreateTeknisi(ctx context.Context, teknisi stock.Teknisi) error
	UpdateTeknisi(ctx context.Context, teknisi stock.Teknisi) error
	DeleteTeknisi(ctx context.Context, id string) error

	GetAllMachines(ctx context.Context) ([]stock.Machine, error)
	GetMachineByID(ctx context.Context, id string) (stock.Machine, error)
	GetMachineByIDCustomer(ctx context.Context, id string) ([]stock.Machine, error)
	CreateMachine(ctx context.Context, machine stock.Machine) error
	UpdateMachine(ctx context.Context, machine stock.Machine) error
	DeleteMachine(ctx context.Context, id string) error

	GetAllCustomers(ctx context.Context) ([]stock.Customer, error)
	GetCustomersFiltered(ctx context.Context, keyword string, page, length int) ([]stock.Customer, int, error)
	CreateCustomer(ctx context.Context, customer stock.Customer) error
	UpdateCustomer(ctx context.Context, customer stock.Customer) error
	DeleteCustomer(ctx context.Context, id string) error

	GetAllRequests(ctx context.Context) ([]stock.Request, error)
	GetRequestsPagination(ctx context.Context, keyword, status string, page, length int) ([]stock.Request, int, error)
	CreateRequest(ctx context.Context, request stock.Request) error
	UpdateRequest(ctx context.Context, request stock.Request) error
	DeleteRequest(ctx context.Context, id string) error

	GetUserByUsername(ctx context.Context, username string) (stock.User, error)
	CreateUser(ctx context.Context, user stock.User) error
	UpdateUser(ctx context.Context, user stock.User) error
	DeleteUser(ctx context.Context, username string) error

	MatchPassword(ctx context.Context, login stock.User) error

	GetAllInventory(ctx context.Context) ([]stock.Inventory, error)
	GetInventoryByID(ctx context.Context, id string) ([]stock.Inventory, error)
	InventoryUsage(ctx context.Context, input stock.InventoryUsage) error
	CreateInventory(ctx context.Context, inventory stock.Inventory) error
	UpdateInventory(ctx context.Context, inventory stock.Inventory) error
	DeleteInventory(ctx context.Context, id_teknisi, id_sparepart string) error

	GetPembelianSparepart(ctx context.Context) ([]stock.PembelianSparepart, error)
	GetPembelianSparepartByID(ctx context.Context, id string) ([]stock.PembelianSparepart, error)
	GetPembelianSparepartBySupplier(ctx context.Context, id string) ([]stock.PembelianSparepart, error)
	CreatePembelianSparepart(ctx context.Context, pembelian_sp stock.PembelianSparepart) error
	UpdatePembelianSparepart(ctx context.Context, pembelian_sp stock.PembelianSparepart) error
	DeletePembelianSparepart(ctx context.Context, id string) error

	GetAllSparepartHistory(ctx context.Context, id_teknisi, id_mesin, id_sparepart string, page, length int) ([]stock.SparepartHistory, int, error)

	GetSuppliersPagination(ctx context.Context, keyword string, page, length int) ([]stock.Supplier, int, error)
	CreateSupplier(ctx context.Context, supplier stock.Supplier) error
	UpdateSupplier(ctx context.Context, supplier stock.Supplier) error
	DeleteSupplier(ctx context.Context, id string) error

	GetAllReturnInventory(ctx context.Context) ([]stock.ReturnInventory, error)
	GetReturnInventoryByStatus(ctx context.Context, status string) ([]stock.ReturnInventory, error)
	ProcessReturnSparepart(ctx context.Context, input stock.ReturnInventory) error
	ApproveReturnInventory(ctx context.Context, input stock.ReturnInventory) error

	GetAllMachineHistories(ctx context.Context, idMachine, idCustomer string, page, length int) ([]stock.MachineHistory, int, error)
	GetMachineHistoryByID(ctx context.Context, idMachine string) ([]stock.MachineHistory, error)
	CreateMachineHistory(ctx context.Context, history stock.MachineHistory) error
	UpdateMachineHistory(ctx context.Context, history stock.MachineHistory) error
	DeleteMachineHistory(ctx context.Context, id int) error
	ReplaceCustomerMachine(ctx context.Context, oldMachineID, newMachineID, customerID string) error
	DeactivateMachine(ctx context.Context, machineID string) error
}

type Handler struct {
	stockSvc StockSvc
}

func New(is StockSvc) *Handler {
	return &Handler{
		stockSvc: is,
	}
}
