package logbook

import "time"

type T_Cont_LogBook struct {
	ID         string    `db:"Log_ID" json:"Log_ID"`
	Date       time.Time `db:"Log_Date" json:"Log_Date"`
	CabAMS     string    `db:"Log_CabAMS" json:"Log_CabAMS"`
	KoliType   string    `db:"Log_KoliType" json:"Log_KoliType"`
	KoliNum    string    `db:"Log_KoliNum" json:"Log_KoliNum"`
	PLNum      string    `db:"Log_PLNum" json:"Log_PLNum"`
	DONum      string    `db:"Log_DONum" json:"Log_DONum"`
	OutDest    string    `db:"Log_OutDest" json:"Log_OutDest"`
	OutName    string    `db:"Log_OutName" json:"Log_OutName"`
	TotBerat   float64   `db:"Log_TotBerat" json:"Log_TotBerat"`
	TotProcod  int64     `db:"Log_TotProcod" json:"Log_TotProcod"`
	LastUpdate time.Time `db:"Log_LastUpdate" json:"Log_LastUpdate"`
	ActiveYN   string    `db:"Log_ActiveYN" json:"Log_ActiveYN"`
}

type M_Cabang_Dist struct {
	DirKode          string    `db:"Dir_kode" json:"Dir_kode"`
	DirNama          string    `db:"Dir_Nama" json:"Dir_Nama"`
	DirAlamat        string    `db:"Dir_alamat" json:"Dir_alamat"`
	DirDaerah        string    `db:"Dir_daerah" json:"Dir_daerah"`
	DirKota          string    `db:"Dir_kota" json:"Dir_kota"`
	DirKDPOS         string    `db:"Dir_KDPOS" json:"Dir_KDPOS"`
	DirTELP1         string    `db:"Dir_TELP1" json:"Dir_TELP1"`
	DirTELP2         string    `db:"Dir_TELP2" json:"Dir_TELP2"`
	DirFAX           string    `db:"Dir_FAX" json:"Dir_FAX"`
	DirEmail         string    `db:"Dir_email" json:"Dir_email"`
	DirUPDATEID      string    `db:"Dir_UPDATEID" json:"Dir_UPDATEID"`
	DirUPDATETIME    time.Time `db:"Dir_UPDATETIME" json:"Dir_UPDATETIME"`
	DirLintang       string    `db:"Dir_Lintang" json:"Dir_Lintang"`
	DirBujur         string    `db:"Dir_Bujur" json:"Dir_Bujur"`
	DirDistributor   string    `db:"Dir_Distributor" json:"Dir_Distributor"`
	DirAktifYN       string    `db:"Dir_AktifYN" json:"Dir_AktifYN"`
	DirKdCabDist     string    `db:"Dir_KdCabDist" json:"Dir_KdCabDist"`
	DirKdDistributor string    `db:"Dir_KdDistributor" json:"Dir_KdDistributor"`
}
