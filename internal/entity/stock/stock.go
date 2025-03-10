package stock

type Sparepart struct {
	ID       string `db:"id_sparepart" json:"id_sparepart"`
	Nama     string `db:"nama_sparepart" json:"nama_sparepart"`
	Quantity int    `db:"quantity" json:"quantity"`
}
