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
	qGetAllSparepartHistory = `SELECT * FROM history_sparepart
								WHERE 
									(id_teknisi LIKE ? OR ? = '')
									AND (id_machine LIKE ? OR ? = '')
									AND (id_sparepart LIKE ? OR ? = '')
								ORDER BY updated_at DESC
								LIMIT ?, ?`

	getAllSparepartHistoryCount  = "GetAllSparepartHistoryCount"
	qGetAllSparepartHistoryCount = `SELECT COUNT(*) FROM history_sparepart
								WHERE 
									(id_teknisi LIKE ? OR ? = '')
									AND (id_machine LIKE ? OR ? = '')
									AND (id_sparepart LIKE ? OR ? = '')`

	getSparepartHistoryByID  = "GetSparepartHistoryByID"
	qGetSparepartHistoryByID = `SELECT * FROM history_sparepart WHERE id_machine = ? ORDER BY updated_at DESC`

	createSparepartHistory  = "CreateSparepartHistory"
	qCreateSparepartHistory = `INSERT INTO history_sparepart(id_history, id_teknisi, id_machine, id_sparepart,
								quantity, counter, updated_by) VALUES (?,?,?,?,?,?,?)`

	updateSparepartHistory  = "UpdateSparepartHistory"
	qUpdateSparepartHistory = `UPDATE history_sparepart 
						SET 
							id_teknisi = COALESCE(NULLIF(?,''), id_teknisi), 
							id_machine = COALESCE(NULLIF(?,''), id_machine), 
							id_sparepart = COALESCE(NULLIF(?,''), id_sparepart), 
							quantity = COALESCE(NULLIF(?,''), quantity), 
							counter = COALESCE(NULLIF(?,''), counter), 
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
	qGetAllMachines = `SELECT id_machine, tipe_machine, COALESCE(id_customer, 'N/A') AS id_customer FROM machine`

	getMachineByID  = "GetMachineByID"
	qGetMachineByID = `SELECT id_machine, tipe_machine, COALESCE(id_customer, 'N/A') AS id_customer FROM machine WHERE id_machine = ?`

	getMachineByIDCustomer  = "GetMachineByIDCustomer"
	qGetMachineByIDCustomer = `SELECT id_machine, tipe_machine, COALESCE(id_customer, 'N/A') AS id_customer FROM machine WHERE id_customer = ?`

	createMachine  = "CreateMachine"
	qCreateMachine = `INSERT INTO machine(id_machine, tipe_machine, id_customer) VALUES (?,?,?)`

	updateMachine  = "UpdateMachine"
	qUpdateMachine = `UPDATE machine 
						SET 
							id_customer = COALESCE(NULLIF(?,''), id_customer)
						WHERE id_machine = ?`

	deleteMachine  = "DeleteMachine"
	qDeleteMachine = "DELETE FROM machine WHERE id_machine = ?"

	// Request
	getAllRequests  = "GetAllRequests"
	qGetAllRequests = `SELECT * FROM request`

	getRequestsPage  = "GetRequestsPage"
	qGetRequestsPage = `SELECT r.*, sp.nama_sparepart, COALESCE(m.tipe_machine, '') AS tipe_machine, t.nama_teknisi, 
						COALESCE(c.nama_customer, '') AS nama_customer FROM request r 
						LEFT JOIN sparepart sp ON r.id_sparepart = sp.id_sparepart 
						LEFT JOIN machine m ON r.id_mesin = m.id_machine
						LEFT JOIN teknisi t ON r.id_teknisi = t.id_teknisi
						LEFT JOIN customer c ON m.id_customer = c.id_customer
						WHERE
							((r.id_teknisi LIKE ? OR ? = '')
						AND
							(r.status_request LIKE ? OR ? = ''))
						ORDER BY updated_at DESC
						LIMIT ?, ?`

	getRequestsCount  = "GetRequestsCount"
	qGetRequestsCount = `SELECT COUNT(*) FROM request WHERE
							((id_teknisi LIKE ? OR ? = '')
						AND
							(status_request LIKE ? OR ? = ''))`

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

	deleteRequestByIDSparepart  = "DeleteRequestByIDSparepart"
	qDeleteRequestByIDSparepart = "DELETE FROM request WHERE id_sparepart = ?"

	// Customer
	getAllCustomers  = "GetAllCustomers"
	qGetAllCustomers = `SELECT * FROM customer`

	getAllCustomersPage  = "GetAllCustomersPage"
	qGetAllCustomersPage = `SELECT * FROM customer WHERE
						((id_customer LIKE ? OR ? = '')
						OR
						(nama_customer LIKE ? OR ? = '')) 
						LIMIT ?, ?`

	getAllCustomersCount  = "GetAllCustomersCount"
	qGetAllCustomersCount = `SELECT COUNT(*) FROM customer WHERE
						((id_customer LIKE ? OR ? = '')
						OR
						(nama_customer LIKE ? OR ? = ''))`

	createCustomer  = "CreateCustomer"
	qCreateCustomer = `INSERT INTO customer(id_customer, company_id, nama_customer, 
						alamat, updated_by) VALUES (?,?,?,?,?)`

	updateCustomer  = "UpdateCustomer"
	qUpdateCustomer = `UPDATE customer 
						SET 
							alamat = COALESCE(NULLIF(?,''), alamat), 
							nama_customer = COALESCE(NULLIF(?,''), nama_customer),
							updated_by = COALESCE(NULLIF(?,''), updated_by)
						WHERE id_customer = ?`

	deleteCustomer  = "DeleteCustomer"
	qDeleteCustomer = "DELETE FROM customer WHERE id_customer = ?"

	//User
	getUserByUsername  = "GetUserByUsername"
	qGetUserByUsername = "SELECT * FROM user WHERE BINARY username = ?"

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

	// Inventory
	getAllInventory  = "GetAllInventory"
	qGetAllInventory = `SELECT i.*, sp.nama_sparepart FROM inventory i LEFT JOIN sparepart sp ON i.id_sparepart = sp.id_sparepart`

	getInventoryByID  = "GetInventoryByID"
	qGetInventoryByID = `SELECT i.*, sp.nama_sparepart FROM inventory i
						LEFT JOIN sparepart sp ON i.id_sparepart = sp.id_sparepart 
						WHERE id_teknisi = ? AND i.quantity > 0`

	getInventoryByIDInv  = "GetInventoryByIDInv"
	qGetInventoryByIDInv = `SELECT i.*, sp.nama_sparepart FROM inventory i
						LEFT JOIN sparepart sp ON i.id_sparepart = sp.id_sparepart 
						WHERE id_inventory = ?`

	createInventory  = "CreateInventory"
	qCreateInventory = `INSERT INTO inventory(id_teknisi, id_sparepart, quantity)
						VALUES (?, ?, ?)
						ON DUPLICATE KEY UPDATE 
							quantity = quantity + VALUES(quantity),
							updated_at = NOW();`

	updateInventory  = "UpdateInventory"
	qUpdateInventory = `UPDATE inventory 
						SET 
							quantity = ?
						WHERE id_teknisi = ? AND id_sparepart = ?`

	deleteInventory  = "DeleteInventory"
	qDeleteInventory = "DELETE FROM inventory WHERE id_teknisi = ? AND id_sparepart = ?"

	//Pembelian Sparepart
	getPembelianSparepart  = "GetPembelianSparepart"
	qGetPembelianSparepart = `SELECT ps.*, s.nama_supplier, sp.nama_sparepart FROM pembelian_sparepart ps 
							JOIN supplier s ON ps.id_supplier = s.id_supplier 
							LEFT JOIN sparepart sp ON ps.id_sparepart = sp.id_sparepart`

	getPembelianSparepartByID  = "GetPembelianSparepartByID"
	qGetPembelianSparepartByID = `SELECT ps.*, s.nama_supplier, sp.nama_sparepart FROM pembelian_sparepart ps 
							JOIN supplier s ON ps.id_supplier = s.id_supplier 
							LEFT JOIN sparepart sp ON ps.id_sparepart = sp.id_sparepart WHERE ps.id_sparepart = ?`

	getPembelianSparepartBySupplier  = "GetPembelianSparepartBySupplier"
	qGetPembelianSparepartBySupplier = `SELECT ps.*, s.nama_supplier, sp.nama_sparepart FROM pembelian_sparepart ps 
							JOIN supplier s ON ps.id_supplier = s.id_supplier 
							LEFT JOIN sparepart sp ON ps.id_sparepart = sp.id_sparepart
							WHERE ps.id_supplier = ?`

	createPembelianSparepart  = "CreatePembelianSparepart"
	qCreatePembelianSparepart = `INSERT INTO pembelian_sparepart(id_sparepart, quantity, harga_per_unit, id_supplier) 
								VALUES (?,?,?,?)`

	updatePembelianSparepart  = "UpdatePembelianSparepart"
	qUpdatePembelianSparepart = `UPDATE pembelian_sparepart
								SET 
									quantity = COALESCE(NULLIF(?,''), quantity),
									harga_per_unit = COALESCE(NULLIF(?,''), harga_per_unit)
								WHERE 
									id_pembelian = ?`

	deletePembelianSparepart  = "DeletePembelianSparepart"
	qDeletePembelianSparepart = `DELETE from pembelian_sparepart WHERE id_pembelian = ?`

	deletePembelianSparepartByIDSparepart  = "DeletePembelianSparepartByIDSparepart"
	qDeletePembelianSparepartByIDSparepart = `DELETE from pembelian_sparepart WHERE id_sparepart = ?`

	getAverageCostSparepart  = "GetAverageCostSparepart"
	qGetAverageCostSparepart = `SELECT COALESCE(SUM(quantity * harga_per_unit) / NULLIF(SUM(quantity), 0), 0) AS average_cost
								FROM pembelian_sparepart WHERE id_sparepart = ?`

	getSuppliers  = "GetSuppliers"
	qGetSuppliers = `SELECT * FROM supplier`

	getSuppliersPage  = "GetSuppliersPage"
	qGetSuppliersPage = `SELECT * FROM supplier WHERE nama_supplier LIKE ? OR ? = '' LIMIT ?, ?`

	getSuppliersCount  = "GetSuppliersCount"
	qGetSuppliersCount = `SELECT COUNT(*) FROM supplier WHERE nama_supplier LIKE ? OR ? = ''`

	createSupplier  = "CreateSupplier"
	qCreateSupplier = `INSERT INTO supplier(nama_supplier) VALUES (?)`

	updateSupplier  = "UpdateSupplier"
	qUpdateSupplier = `UPDATE supplier
					SET
						nama_supplier = COALESCE(NULLIF(?,''), nama_supplier)
					WHERE 
						id_supplier = ?`

	deleteSupplier  = "DeleteSupplier"
	qDeleteSupplier = `DELETE FROM supplier WHERE id_supplier = ?`

	getAllReturnInventory  = "GetAllReturnInventory"
	qGetAllReturnInventory = `SELECT ri.*, sp.nama_sparepart, t.nama_teknisi 
							FROM return_inventory ri
							LEFT JOIN sparepart sp ON ri.id_sparepart = sp.id_sparepart
							JOIN teknisi t ON ri.returned_by = t.id_teknisi
							ORDER BY returned_at DESC`

	getReturnInventoryByStatus  = "GetReturnInventoryByStatus"
	qGetReturnInventoryByStatus = `SELECT ri.*, sp.nama_sparepart 
								FROM return_inventory ri
								LEFT JOIN sparepart sp ON ri.id_sparepart = sp.id_sparepart
								WHERE ri.status = ?
								ORDER BY returned_at DESC`

	createReturnInventory  = "CreateReturnInventory"
	qCreateReturnInventory = `INSERT INTO return_inventory(id_inventory, id_sparepart, quantity, returned_by) VALUES (?, ?, ?, ?)`

	approveReturnInventory  = "ApproveReturnInventory"
	qApproveReturnInventory = `UPDATE return_inventory
							SET status = ?, approved_by = ?, approved_at = NOW()
							WHERE id_return = ?`

	// Machine History
	getAllMachineHistory  = "GetAllMachineHistory"
	qGetAllMachineHistory = `SELECT * FROM machine_history 
                         WHERE 
                           (id_machine LIKE ? OR ? = '') AND 
                           (id_customer LIKE ? OR ? = '') 
                         ORDER BY tanggal_mulai DESC 
                         LIMIT ?, ?`

	getAllMachineHistoryCount  = "GetAllMachineHistoryCount"
	qGetAllMachineHistoryCount = `SELECT COUNT(*) FROM machine_history 
                              WHERE 
                                (id_machine LIKE ? OR ? = '') AND 
                                (id_customer LIKE ? OR ? = '')`

	getMachineHistoryByID  = "GetMachineHistoryByID"
	qGetMachineHistoryByID = `SELECT mh.*, c.nama_customer FROM machine_history mh
							JOIN customer c ON mh.id_customer = c.id_customer
                        	WHERE id_machine = ? 
                        	ORDER BY tanggal_mulai DESC`

	getLastMachineHistoryByMachineID  = "GetLastMachineHistoryByMachineID"
	qGetLastMachineHistoryByMachineID = `SELECT id_history, id_machine, id_customer, status, tanggal_mulai, tanggal_selesai
										FROM machine_history
										WHERE id_machine = ? AND tanggal_mulai <= NOW()
										ORDER BY tanggal_mulai DESC
										LIMIT 1`

	createMachineHistory  = "CreateMachineHistory"
	qCreateMachineHistory = `INSERT INTO machine_history(id_machine, id_customer, tanggal_mulai, status) 
                         VALUES (?, ?, ?, ?)`

	updateMachineHistory  = "UpdateMachineHistory"
	qUpdateMachineHistory = `UPDATE machine_history 
                         SET 
                           id_customer = COALESCE(NULLIF(?,''), id_customer),
                           tanggal_mulai = COALESCE(NULLIF(?,''), tanggal_mulai),
                           tanggal_selesai = COALESCE(NULLIF(?,''), tanggal_selesai),
                           status = COALESCE(NULLIF(?,''), status)
                         WHERE id_history = ?`

	deleteMachineHistory  = "DeleteMachineHistory"
	qDeleteMachineHistory = `DELETE FROM machine_history WHERE id_history = ?`
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
		{getAllSparepartHistoryCount, qGetAllSparepartHistoryCount},
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
		{getAllCustomersPage, qGetAllCustomersPage},
		{getAllCustomersCount, qGetAllCustomersCount},
		//user
		{getUserByUsername, qGetUserByUsername},
		//inventory
		{getAllInventory, qGetAllInventory},
		{getInventoryByID, qGetInventoryByID},
		{getInventoryByIDInv, qGetInventoryByIDInv},
		//pembelian sparepart
		{getPembelianSparepart, qGetPembelianSparepart},
		{getPembelianSparepartByID, qGetPembelianSparepartByID},
		{getPembelianSparepartBySupplier, qGetPembelianSparepartBySupplier},
		{getAverageCostSparepart, qGetAverageCostSparepart},
		//supplier
		{getSuppliers, qGetSuppliers},
		{getSuppliersPage, qGetSuppliersPage},
		{getSuppliersCount, qGetSuppliersCount},
		//return sparepart
		{getAllReturnInventory, qGetAllReturnInventory},
		{getReturnInventoryByStatus, qGetReturnInventoryByStatus},
		//machine history
		{getAllMachineHistory, qGetAllMachineHistory},
		{getAllMachineHistoryCount, qGetAllMachineHistoryCount},
		{getMachineHistoryByID, qGetMachineHistoryByID},
		{getLastMachineHistoryByMachineID, qGetLastMachineHistoryByMachineID},
	}
	insertStmt = []statement{
		{createSparepart, qCreateSparepart},
		{createSparepartHistory, qCreateSparepartHistory},
		{createTeknisi, qCreateTeknisi},
		{createMachine, qCreateMachine},
		{createRequest, qCreateRequest},
		{createCustomer, qCreateCustomer},
		{createUser, qCreateUser},
		{createInventory, qCreateInventory},
		{createPembelianSparepart, qCreatePembelianSparepart},
		{createSupplier, qCreateSupplier},
		{createReturnInventory, qCreateReturnInventory},
		{createMachineHistory, qCreateMachineHistory},
	}
	updateStmt = []statement{
		{updateSparepart, qUpdateSparepart},
		{updateSparepartHistory, qUpdateSparepartHistory},
		{updateTeknisi, qUpdateTeknisi},
		{updateMachine, qUpdateMachine},
		{updateRequest, qUpdateRequest},
		{updateCustomer, qUpdateCustomer},
		{updateUser, qUpdateUser},
		{updateInventory, qUpdateInventory},
		{updatePembelianSparepart, qUpdatePembelianSparepart},
		{updateSupplier, qUpdateSupplier},
		{approveReturnInventory, qApproveReturnInventory},
		{updateMachineHistory, qUpdateMachineHistory},
	}
	deleteStmt = []statement{
		{deleteSparepart, qDeleteSparepart},
		{deleteSparepartHistory, qDeleteSparepartHistory},
		{deleteTeknisi, qDeleteTeknisi},
		{deleteMachine, qDeleteMachine},
		{deleteRequest, qDeleteRequest},
		{deleteRequestByIDSparepart, qDeleteRequestByIDSparepart},
		{deleteCustomer, qDeleteCustomer},
		{deleteUser, qDeleteUser},
		{deleteInventory, qDeleteInventory},
		{deletePembelianSparepart, qDeletePembelianSparepart},
		{deletePembelianSparepartByIDSparepart, qDeletePembelianSparepartByIDSparepart},
		{deleteSupplier, qDeleteSupplier},
		{deleteMachineHistory, qDeleteMachineHistory},
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
