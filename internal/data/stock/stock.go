package stock

import (
	"context"
	"log"

	"github.com/jmoiron/sqlx"
)

type (
	// Data ...
	Data struct {
		db   *sqlx.DB
		stmt map[string]*sqlx.Stmt
	}

	// Key - value map for query selector
	statement struct {
		key   string
		query string
	}
)

// Query constants
const (
	// Sparepart
	getAllSpareparts  = "GetAllSpareparts"
	qGetAllSpareparts = `SELECT * FROM sparepart`

	getAllSparepartsPage  = "GetAllSparepartsPage"
	qGetAllSparepartsPage = `SELECT * FROM sparepart WHERE 
								((id_sparepart LIKE ? OR ? = '')
							OR
								(nama_sparepart LIKE ? OR ? = '')) 
							LIMIT ?, ?`

	getAllSparepartsCount  = "GetAllSparepartsCount"
	qGetAllSparepartsCount = `SELECT COUNT(*) FROM sparepart WHERE 
								((id_sparepart LIKE ? OR ? = '')
							OR
								(nama_sparepart LIKE ? OR ? = ''))`

	getSparepartByID  = "GetSparepartByID"
	qGetSparepartByID = `SELECT * FROM sparepart WHERE id_sparepart = ?`

	createSparepart  = "CreateSparepart"
	qCreateSparepart = `INSERT INTO sparepart(id_sparepart, nama_sparepart, quantity) VALUES (?,?,?)`

	updateSparepart  = "UpdateSparepart"
	qUpdateSparepart = `UPDATE sparepart 
						SET 
							nama_sparepart = COALESCE(NULLIF(?,''), nama_sparepart), 
							quantity = COALESCE(NULLIF(?,''), quantity)
						WHERE id_sparepart = ?`

	deleteSparepart  = "DeleteSparepart"
	qDeleteSparepart = "DELETE FROM sparepart WHERE id_sparepart = ?"

	//History Sparepart
	getAllSparepartHistory  = "GetAllSparepartHistory"
	qGetAllSparepartHistory = `SELECT * FROM history_sparepart`

	getSparepartHistoryByID  = "GetSparepartHistoryByID"
	qGetSparepartHistoryByID = `SELECT * FROM history_sparepart WHERE id_machine = ?`

	createSparepartHistory  = "CreateSparepartHistory"
	qCreateSparepartHistory = `INSERT INTO history_sparepart(id_history, id_teknisi, id_machine, id_sparepart, id_request,
								quantity, updated_by) VALUES (?,?,?,?,?,?,?)`

	updateSparepartHistory  = "UpdateSparepartHistory"
	qUpdateSparepartHistory = `UPDATE history_sparepart 
						SET 
							id_teknisi = COALESCE(NULLIF(?,''), id_teknisi), 
							id_machine = COALESCE(NULLIF(?,''), id_machine), 
							id_sparepart = COALESCE(NULLIF(?,''), id_sparepart), 
							id_request = COALESCE(NULLIF(?,''), id_request), 
							quantity = COALESCE(NULLIF(?,''), quantity), 
							updated_by = COALESCE(NULLIF(?,''), updated_by),
							updated_at = COALESCE(NULLIF(?,''), updated_at)
						WHERE id_history = ?`

	deleteSparepartHistory  = "DeleteSparepartHistory"
	qDeleteSparepartHistory = "DELETE FROM history_sparepart WHERE id_history = ?"

	//Teknisi
	getAllTeknisi  = "GetAllTeknisi"
	qGetAllTeknisi = `SELECT * FROM teknisi`

	getTeknisiByID  = "GetTeknisiByID"
	qGetTeknisiByID = `SELECT * FROM teknisi WHERE id_teknisi = ?`

	createTeknisi  = "CreateTeknisi"
	qCreateTeknisi = `INSERT INTO teknisi(id_teknisi, nama_teknisi, active_yn) VALUES (?,?,?)`

	updateTeknisi  = "UpdateTeknisi"
	qUpdateTeknisi = `UPDATE teknisi 
						SET 
							nama_teknisi = COALESCE(NULLIF(?,''), nama_teknisi), 
							active_yn = COALESCE(NULLIF(?,''), active_yn)
						WHERE id_teknisi = ?`

	deleteTeknisi  = "DeleteTeknisi"
	qDeleteTeknisi = "DELETE FROM teknisi WHERE id_teknisi = ?"

	//Mesin
	getAllMachines  = "GetAllMachines"
	qGetAllMachines = `SELECT * FROM machine`

	getMachineByID  = "GetMachineByID"
	qGetMachineByID = `SELECT * FROM machine WHERE id_machine = ?`

	getMachineByIDCustomer  = "GetMachineByIDCustomer"
	qGetMachineByIDCustomer = `SELECT * FROM machine WHERE id_customer = ?`

	createMachine  = "CreateMachine"
	qCreateMachine = `INSERT INTO machine(id_machine, tipe_machine, counter, id_customer) VALUES (?,?,?,?)`

	updateMachine  = "UpdateMachine"
	qUpdateMachine = `UPDATE machine 
						SET 
							counter = COALESCE(NULLIF(?,''), counter), 
							id_customer = COALESCE(NULLIF(?,''), id_customer)
						WHERE id_machine = ?`

	deleteMachine  = "DeleteMachine"
	qDeleteMachine = "DELETE FROM machine WHERE id_machine = ?"

	// Request
	getAllRequests  = "GetAllRequests"
	qGetAllRequests = `SELECT * FROM request`

	getRequestsPage  = "GetRequestsPage"
	qGetRequestsPage = `SELECT r.*, sp.nama_sparepart, m.tipe_machine, t.nama_teknisi, c.nama_customer FROM request r 
						JOIN sparepart sp ON r.id_sparepart = sp.id_sparepart 
						JOIN machine m ON r.id_mesin = m.id_machine
						JOIN teknisi t ON r.id_teknisi = t.id_teknisi
						JOIN customer c ON m.id_customer = c.id_customer
						LIMIT ?, ?`

	getRequestsCount  = "GetRequestsCount"
	qGetRequestsCount = `SELECT COUNT(*) FROM request`

	createRequest  = "CreateRequest"
	qCreateRequest = `INSERT INTO request(id_teknisi, id_mesin, 
						id_sparepart, quantity, status_request, updated_by) VALUES (?,?,?,?,?,?)`

	updateRequest  = "UpdateRequest"
	qUpdateRequest = `UPDATE request 
						SET 
							id_mesin = COALESCE(NULLIF(?,''), id_mesin), 
							id_sparepart = COALESCE(NULLIF(?,''), id_sparepart),
							quantity = COALESCE(NULLIF(?,''), quantity), 
							status_request = COALESCE(NULLIF(?,''), status_request),
							updated_by = COALESCE(NULLIF(?,''), updated_by)
						WHERE id_request = ?`

	deleteRequest  = "DeleteRequest"
	qDeleteRequest = "DELETE FROM request WHERE id_request = ?"

	// Customer
	getAllCustomers  = "GetAllCustomers"
	qGetAllCustomers = `SELECT * FROM customer`

	createCustomer  = "CreateCustomer"
	qCreateCustomer = `INSERT INTO customer(id_customer, company_id, nama_customer, 
						alamat, pic, updated_by, updated_at) VALUES (?,?,?,?,?,?,?)`

	updateCustomer  = "UpdateCustomer"
	qUpdateCustomer = `UPDATE customer 
						SET 
							alamat = COALESCE(NULLIF(?,''), alamat), 
							nama_customer = COALESCE(NULLIF(?,''), nama_customer),
							pic = COALESCE(NULLIF(?,''), pic), 
							updated_by = COALESCE(NULLIF(?,''), updated_by), 
							updated_at = COALESCE(NULLIF(?,''), updated_at)
						WHERE id_customer = ?`

	deleteCustomer  = "DeleteCustomer"
	qDeleteCustomer = "DELETE FROM customer WHERE id_customer = ?"

	//User
	getUserByUsername  = "GetUserByUsername"
	qGetUserByUsername = "SELECT * FROM user WHERE username = ?"

	createUser  = "CreateUser"
	qCreateUser = `INSERT INTO user(username, password) VALUES (?,?)`

	updateUser  = "UpdateUser"
	qUpdateUser = `UPDATE user
					SET 
						password = COALESCE(NULLIF(?,''), password)
					WHERE 
						username = ?`

	deleteUser  = "DeleteUser"
	qDeleteUser = `DELETE from user WHERE username = ?`
)

var (
	readStmt = []statement{
		//sparepart
		{getAllSpareparts, qGetAllSpareparts},
		{getAllSparepartsPage, qGetAllSparepartsPage},
		{getAllSparepartsCount, qGetAllSparepartsCount},
		{getSparepartByID, qGetSparepartByID},
		//history
		{getAllSparepartHistory, qGetAllSparepartHistory},
		{getSparepartHistoryByID, qGetSparepartHistoryByID},
		//teknisi
		{getAllTeknisi, qGetAllTeknisi},
		{getTeknisiByID, qGetTeknisiByID},
		//machine
		{getAllMachines, qGetAllMachines},
		{getMachineByID, qGetMachineByID},
		{getMachineByIDCustomer, qGetMachineByIDCustomer},
		//request
		{getAllRequests, qGetAllRequests},
		{getRequestsPage, qGetRequestsPage},
		{getRequestsCount, qGetRequestsCount},
		//customer
		{getAllCustomers, qGetAllCustomers},
		//user
		{getUserByUsername, qGetUserByUsername},
	}
	insertStmt = []statement{
		{createSparepart, qCreateSparepart},
		{createSparepartHistory, qCreateSparepartHistory},
		{createTeknisi, qCreateTeknisi},
		{createMachine, qCreateMachine},
		{createRequest, qCreateRequest},
		{createCustomer, qCreateCustomer},
		{createUser, qCreateUser},
	}
	updateStmt = []statement{
		{updateSparepart, qUpdateSparepart},
		{updateSparepartHistory, qUpdateSparepartHistory},
		{updateTeknisi, qUpdateTeknisi},
		{updateMachine, qUpdateMachine},
		{updateRequest, qUpdateRequest},
		{updateCustomer, qUpdateCustomer},
		{updateUser, qUpdateUser},
	}
	deleteStmt = []statement{
		{deleteSparepart, qDeleteSparepart},
		{deleteSparepartHistory, qDeleteSparepartHistory},
		{deleteTeknisi, qDeleteTeknisi},
		{deleteMachine, qDeleteMachine},
		{deleteRequest, qDeleteRequest},
		{deleteCustomer, qDeleteCustomer},
		{deleteUser, qDeleteUser},
	}
)

// New returns data instance
func New(db *sqlx.DB) Data {
	d := Data{
		db: db,
	}

	d.initStmt()
	return d
}

// Initialize queries as prepared statements
func (d *Data) initStmt() {
	var (
		err   error
		stmts = make(map[string]*sqlx.Stmt)
	)

	for _, v := range readStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize statement key %v, err : %v", v.key, err)
		}
	}
	for _, v := range insertStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize insert statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range updateStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize update statement key %v, err : %v", v.key, err)
		}
	}

	for _, v := range deleteStmt {
		stmts[v.key], err = d.db.PreparexContext(context.Background(), v.query)
		if err != nil {
			log.Fatalf("[DB] Failed to initialize delete statement key %v, err : %v", v.key, err)
		}
	}

	d.stmt = stmts
}
