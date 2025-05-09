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
	ID       string `db:"id_machine" json:"id_machine"`
	Type     string `db:"tipe_machine" json:"tipe_machine"`
	Counter  int    `db:"counter" json:"counter"`
	Customer string `db:"id_customer" json:"id_customer"`
	History  []SparepartHistory
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
	Status         string    `db:"status_request" json:"status_request"`
	TanggalRequest time.Time `db:"tanggal_request" json:"tanggal_request"`
	UpdatedBy      string    `db:"updated_by" json:"updated_by"`
	UpdatedAt      time.Time `db:"updated_at" json:"updated_at"`
	Customer       string    `db:"nama_customer" json:"nama_customer"`
}

type SparepartHistory struct {
	IDHistory     int       `db:"id_history" json:"id_history"`
	IDTeknisi     string    `db:"id_teknisi" json:"id_teknisi"`
	NamaTeknisi   string    `json:"nama_teknisi"`
	IDMachine     string    `db:"id_machine" json:"id_machine"`
	IDSparepart   string    `db:"id_sparepart" json:"id_sparepart"`
	NamaSparepart string    `json:"nama_sparepart"`
	IDRequest     int       `db:"id_request" json:"id_request"`
	Quantity      int       `db:"quantity" json:"quantity"`
	Counter       int       `db:"counter" json:"counter"`
	UpdatedBy     string    `db:"updated_by" json:"updated_by"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}

type Customer struct {
	ID           string    `db:"id_customer" json:"id_customer"`
	Company      int       `db:"company_id" json:"company_id"`
	NamaCustomer string    `db:"nama_customer" json:"nama_customer"`
	Alamat       string    `db:"alamat" json:"alamat"`
	PIC          string    `db:"pic" json:"pic"`
	UpdatedBy    string    `db:"updated_by" json:"updated_by"`
	UpdatedAt    time.Time `db:"updated_at" json:"updated_at"`
}

type Inventory struct {
	ID            string    `db:"id_inventory" json:"id_inventory"`
	Teknisi       string    `db:"id_teknisi" json:"id_teknisi"`
	Sparepart     string    `db:"id_sparepart" json:"id_sparepart"`
	NamaSparepart string    `db:"nama_sparepart" json:"nama_sparepart"`
	Quantity      int       `db:"quantity" json:"quantity"`
	UpdatedAt     time.Time `db:"updated_at" json:"updated_at"`
}

type PembelianSparepart struct {
	ID               string    `db:"id_pembelian" json:"id_pembelian"`
	Sparepart        string    `db:"id_sparepart" json:"id_sparepart"`
	TanggalPembelian time.Time `db:"tanggal_pembelian" json:"tanggal_pembelian"`
	Quantity         int       `db:"quantity" json:"quantity"`
	HargaPerUnit     float64   `db:"harga_per_unit" json:"harga_per_unit"`
}

type SparepartCostFilter struct {
	Year       *int
	Month      *int
	CustomerID *string
	MachineID  *string
}

type SparepartCostResult struct {
	CustomerID   string  `db:"id_customer"`   // Memetakan ke kolom id_customer
	CustomerName string  `db:"nama_customer"` // Memetakan ke kolom nama_customer
	MachineID    string  `db:"id_machine"`    // Memetakan ke kolom id_machine
	Month        string  `db:"bulan"`         // Memetakan ke kolom bulan
	Year         int     `db:"tahun"`         // Memetakan ke kolom tahun
	TotalCost    float64 `db:"total_cost"`    // Memetakan ke kolom total_cost
}

type User struct {
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
}
