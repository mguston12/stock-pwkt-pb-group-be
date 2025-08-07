package stock

import "time"

type Sparepart struct {
	ID          string  `db:"id_sparepart" json:"id_sparepart"`
	Nama        string  `db:"nama_sparepart" json:"nama_sparepart"`
	Quantity    int     `db:"quantity" json:"quantity"`
	AverageCost float64 `json:"average_cost"`
}

type Teknisi struct {
	ID       string `db:"id_teknisi" json:"id_teknisi"`
	Nama     string `db:"nama_teknisi" json:"nama_teknisi"`
	ActiveYN string `db:"active_yn" json:"active_yn"`
}

type Machine struct {
	ID           string `db:"id_machine" json:"id_machine"`
	Type         string `db:"tipe_machine" json:"tipe_machine"`
	Counter      int    `db:"counter" json:"counter"`
	Customer     string `db:"id_customer" json:"id_customer"`
	NamaCustomer string `db:"nama_customer" json:"nama_customer"`
	Alamat       string `db:"alamat" json:"alamat"`
	SerialNumber string `db:"serial_number" json:"serial_number"`
	History      []SparepartHistory
}

type Request struct {
	ID             int       `db:"id_request" json:"id_request"`
	Teknisi        string    `db:"id_teknisi" json:"id_teknisi"`
	NamaTeknisi    string    `db:"nama_teknisi" json:"nama_teknisi"`
	Mesin          string    `db:"id_mesin" json:"id_mesin"`
	TipeMesin      string    `db:"tipe_machine" json:"tipe_machine"`
	Sparepart      string    `db:"id_sparepart" json:"id_sparepart"`
	NamaSparepart  string    `db:"nama_sparepart" json:"nama_sparepart"`
	Quantity       int       `db:"quantity" json:"quantity"`
	StockQuantity  int       `db:"stock_quantity" json:"stock_quantity"`
	Status         string    `db:"status_request" json:"status_request"`
	TanggalRequest time.Time `db:"tanggal_request" json:"tanggal_request"`
	UpdatedBy      string    `db:"updated_by" json:"updated_by"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
	Customer       string    `db:"nama_customer" json:"nama_customer"`
}

type SparepartHistory struct {
	IDHistory       int       `db:"id_history" json:"id_history"`
	IDTeknisi       string    `db:"id_teknisi" json:"id_teknisi"`
	NamaTeknisi     string    `json:"nama_teknisi"`
	IDMachine       string    `db:"id_machine" json:"id_machine"`
	IDSparepart     string    `db:"id_sparepart" json:"id_sparepart"`
	NamaSparepart   string    `json:"nama_sparepart"`
	Quantity        int       `db:"quantity" json:"quantity"`
	Counter         int       `db:"counter" json:"counter"`
	CounterColour   int       `db:"counter_colour" json:"counter_colour"`
	CounterColourA3 int       `db:"counter_colour_a3" json:"counter_colour_a3"`
	UpdatedBy       string    `db:"updated_by" json:"updated_by"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}

type ReportData struct {
	IDHistory     int       `db:"id_history" json:"id_history"`
	IDTeknisi     string    `db:"id_teknisi" json:"id_teknisi"`
	NamaTeknisi   string    `db:"nama_teknisi" json:"nama_teknisi"`
	IDMachine     string    `db:"id_machine" json:"id_machine"`
	IDSparepart   string    `db:"id_sparepart" json:"id_sparepart"`
	NamaSparepart string    `db:"nama_sparepart" json:"nama_sparepart"`
	NamaCustomer  string    `db:"nama_customer" json:"nama_customer"`
	Quantity      int       `db:"quantity" json:"quantity"`
	Counter       int       `db:"counter" json:"counter"`
	AverageCost   float64   `db:"average_cost" json:"average_cost"`
	UpdatedBy     string    `db:"updated_by" json:"updated_by"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}

type Customer struct {
	ID           string    `db:"id_customer" json:"id_customer"`
	Company      int       `db:"company_id" json:"company_id"`
	NamaCustomer string    `db:"nama_customer" json:"nama_customer"`
	Alamat       string    `db:"alamat" json:"alamat"`
	UpdatedBy    string    `db:"updated_by" json:"updated_by"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

type Inventory struct {
	ID            int       `db:"id_inventory" json:"id_inventory"`
	Teknisi       string    `db:"id_teknisi" json:"id_teknisi"`
	Sparepart     string    `db:"id_sparepart" json:"id_sparepart"`
	NamaSparepart string    `db:"nama_sparepart" json:"nama_sparepart"`
	Quantity      int       `db:"quantity" json:"quantity"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}

type PembelianSparepart struct {
	ID               string    `db:"id_pembelian" json:"id_pembelian"`
	Sparepart        string    `db:"id_sparepart" json:"id_sparepart"`
	NamaSparepart    string    `db:"nama_sparepart" json:"nama_sparepart"`
	TanggalPembelian time.Time `db:"tanggal_pembelian" json:"tanggal_pembelian"`
	Quantity         int       `db:"quantity" json:"quantity"`
	HargaPerUnit     float64   `db:"harga_per_unit" json:"harga_per_unit"`
	Supplier         int       `db:"id_supplier" json:"id_supplier"`
	NamaSupplier     string    `db:"nama_supplier" json:"nama_supplier"`
}

type ReturnInventory struct {
	IDReturn      int        `db:"id_return" json:"id_return"`
	IDInventory   int        `db:"id_inventory" json:"id_inventory"`
	IDSparepart   string     `db:"id_sparepart" json:"id_sparepart"`
	Teknisi       string     `db:"nama_teknisi" json:"nama_teknisi"`
	Quantity      int        `db:"quantity" json:"quantity"`
	ReturnedBy    string     `db:"returned_by" json:"returned_by"`
	ReturnedAt    time.Time  `db:"returned_at" json:"returned_at"`
	Status        string     `db:"status" json:"status"`
	ApprovedBy    *string    `db:"approved_by" json:"approved_by"`
	ApprovedAt    *time.Time `db:"approved_at" json:"approved_at"`
	NamaSparepart string     `db:"nama_sparepart" json:"nama_sparepart"`
}

type MachineHistory struct {
	IDHistory          int        `db:"id_history" json:"id_history"`
	IDMachine          string     `db:"id_machine" json:"id_machine"`
	IDCustomer         string     `db:"id_customer" json:"id_customer"`
	NamaCustomer       string     `db:"nama_customer" json:"nama_customer"`
	TanggalMulai       time.Time  `db:"tanggal_mulai" json:"tanggal_mulai"`
	TanggalMulaiString string     `db:"tanggal_mulai_string" json:"tanggal_mulai_string"`
	TanggalSelesai     *time.Time `db:"tanggal_selesai" json:"tanggal_selesai"`
	Status             string     `db:"status" json:"status"`
}

type InventoryUsage struct {
	InventoryID     int    `json:"id_inventory"`
	MachineID       string `json:"id_machine"`
	Quantity        int    `json:"quantity"`
	Counter         int    `json:"counter"`
	CounterColour   int    `json:"counter_colour"`
	CounterColourA3 int    `json:"counter_colour_a3"`
	UpdatedBy       string `json:"updated_by"`
}

type SparepartCostFilter struct {
	Year       *int
	Month      *int
	CustomerID *string
	MachineID  *string
}

type SparepartCostResult struct {
	CustomerID   string  `db:"id_customer"`
	CustomerName string  `db:"nama_customer"`
	MachineID    string  `db:"id_machine"`
	Month        string  `db:"bulan"`
	Year         int     `db:"tahun"`
	TotalCost    float64 `db:"total_cost"`
}

type User struct {
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
	Role     string `db:"role" json:"role"`
}

type Supplier struct {
	ID   int    `db:"id_supplier" json:"id_supplier"`
	Nama string `db:"nama_supplier" json:"nama_supplier"`
}

type LoginResponse struct {
	Message string `json:"message"`
}
