package stock

import (
	"context"
	"database/sql"
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

	GetAllSparepartHistory(ctx context.Context, id_teknisi, id_mesin, id_sparepart string, offset, limit int) ([]stock.SparepartHistory, error)
	GetAllSparepartHistoryCount(ctx context.Context, id_teknisi, id_mesin, id_sparepart string) ([]stock.SparepartHistory, int, error)
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
	GetAllMachinesPage(ctx context.Context, keyword string, offset, limit int) ([]stock.Machine, error)
	GetAllMachinesCount(ctx context.Context, keyword string) ([]stock.Machine, int, error)
	GetMachineByID(ctx context.Context, id string) (stock.Machine, error)
	GetMachineByIDCustomer(ctx context.Context, id string) ([]stock.Machine, error)
	CreateMachine(ctx context.Context, machine stock.Machine) error
	UpdateMachine(ctx context.Context, machine stock.Machine) error
	DeleteMachine(ctx context.Context, id string) error

	GetAllCustomers(ctx context.Context) ([]stock.Customer, error)
	GetAllCustomersPage(ctx context.Context, keyword string, offset, limit int) ([]stock.Customer, error)
	GetAllCustomersCount(ctx context.Context, keyword string) ([]stock.Customer, int, error)
	CreateCustomer(ctx context.Context, customer stock.Customer) error
	UpdateCustomer(ctx context.Context, customer stock.Customer) error
	DeleteCustomer(ctx context.Context, id string) error

	FetchAndIncreaseCounter(ctx context.Context, company int) (int, error)

	GetAllRequests(ctx context.Context) ([]stock.Request, error)
	GetRequestsPage(ctx context.Context, teknisi, status string, offset, limit int) ([]stock.Request, error)
	GetRequestsCount(ctx context.Context, teknisi, status string) ([]stock.Request, int, error)
	CreateRequest(ctx context.Context, request stock.Request) error
	CreateRequestTx(ctx context.Context, tx *sql.Tx, request stock.Request) error

	UpdateRequest(ctx context.Context, request stock.Request) error
	DeleteRequest(ctx context.Context, id string) error
	DeleteRequestByIDSparepart(ctx context.Context, id string) error

	GetUserByUsername(ctx context.Context, username string) (stock.User, error)
	CreateUser(ctx context.Context, user stock.User) error
	UpdateUser(ctx context.Context, user stock.User) error
	DeleteUser(ctx context.Context, username string) error

	GetAllInventory(ctx context.Context) ([]stock.Inventory, error)
	GetInventoryByID(ctx context.Context, id string) ([]stock.Inventory, error)
	GetInventoryByIDInv(ctx context.Context, id int) (stock.Inventory, error)
	CreateInventory(ctx context.Context, inventory stock.Inventory) error
	UpdateInventory(ctx context.Context, inventory stock.Inventory) error
	DeleteInventory(ctx context.Context, id_teknisi, id_sparepart string) error

	GetPembelianSparepart(ctx context.Context) ([]stock.PembelianSparepart, error)
	GetPembelianSparepartByID(ctx context.Context, id string) ([]stock.PembelianSparepart, error)
	GetPembelianSparepartBySupplier(ctx context.Context, id string) ([]stock.PembelianSparepart, error)
	CreatePembelianSparepart(ctx context.Context, pembelian_sp stock.PembelianSparepart) error
	UpdatePembelianSparepart(ctx context.Context, pembelian_sp stock.PembelianSparepart) error
	DeletePembelianSparepart(ctx context.Context, id string) error
	DeletePembelianSparepartByIDSparepart(ctx context.Context, id string) error
	GetAverageCostSparepart(ctx context.Context, id string) (float64, error)
	GetSparepartCost(ctx context.Context, filter stock.SparepartCostFilter) ([]stock.SparepartCostResult, error)

	GetSuppliers(ctx context.Context) ([]stock.Supplier, error)
	GetSuppliersPage(ctx context.Context, nama string, offset, limit int) ([]stock.Supplier, error)
	GetSuppliersCount(ctx context.Context, nama string) ([]stock.Supplier, int, error)
	CreateSupplier(ctx context.Context, supplier stock.Supplier) error
	UpdateSupplier(ctx context.Context, supplier stock.Supplier) error
	DeleteSupplier(ctx context.Context, id string) error

	CheckSparepartValidOrNot(ctx context.Context, id string) (int, error)

	GetAllReturnInventory(ctx context.Context) ([]stock.ReturnInventory, error)
	GetReturnInventoryByStatus(ctx context.Context, status string) ([]stock.ReturnInventory, error)
	CreateReturnInventory(ctx context.Context, input stock.ReturnInventory) error
	ApproveReturnInventory(ctx context.Context, input stock.ReturnInventory) error

	GetAllMachineHistory(ctx context.Context, idMachine, idCustomer string, offset, limit int) ([]stock.MachineHistory, error)
	GetMachineHistoryByID(ctx context.Context, idMachine string) ([]stock.MachineHistory, error)
	GetAllMachineHistoryCount(ctx context.Context, idMachine, idCustomer string) (int, error)
	GetLastMachineHistoryByMachineID(ctx context.Context, id string) (stock.MachineHistory, error)
	CreateMachineHistory(ctx context.Context, history stock.MachineHistory) error
	UpdateMachineHistory(ctx context.Context, history stock.MachineHistory) error
	DeleteMachineHistory(ctx context.Context, id int) error

	GetSparepartHistoryByMonth(ctx context.Context, bulan, tahun int) ([]stock.ReportData, error)

	BeginTx(ctx context.Context) (*sql.Tx, error)

	GetInventoryByIDInvTx(ctx context.Context, tx *sql.Tx, id int) (stock.Inventory, error)
	UpdateInventoryTx(ctx context.Context, tx *sql.Tx, inv stock.Inventory) error
	CreateSparepartHistoryTx(ctx context.Context, tx *sql.Tx, history stock.SparepartHistory) error

	GetVisitsByID(ctx context.Context, id int) ([]stock.Visit, error)
	CreateVisit(ctx context.Context, visit stock.Visit) error
	UpdateVisit(ctx context.Context, visit stock.Visit) error
	DeleteVisit(ctx context.Context, id int) error
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
