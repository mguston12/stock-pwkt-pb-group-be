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
	getAllSpareparts  = "GetAllSpareparts"
	qGetAllSpareparts = `SELECT * FROM sparepart`

	createSparepart  = "CreateSparepart"
	qCreateSparepart = `INSERT INTO sparepart(id_sparepart, nama_sparepart, quantity) VALUES (?,?,?)`

	updateSparepart  = "UpdateSparepart"
	qUpdateSparepart = `UPDATE sparepart 
						SET 
							nama_sparepart = COALESCE(NULLIF(?,''), nama_sparepart), 
							quantity = COALESCE(NULLIF(?,''), quantity)
						WHERE id_sparepart = ?`

	deleteSparepart  = "DeleteSparepart"
	qDeleteSparepart = "DELETE FROM sparepart where id_sparepart = ?"

	getAllTeknisi  = "GetAllTeknisi"
	qGetAllTeknisi = `SELECT * FROM teknisi`

	createTeknisi  = "CreateTeknisi"
	qCreateTeknisi = `INSERT INTO sparepart(id_teknisi, nama_teknisi, active_yn) VALUES (?,?,?)`

	updateTeknisi  = "UpdateTeknisi"
	qUpdateTeknisi = `UPDATE teknisi 
						SET 
							nama_teknisi = COALESCE(NULLIF(?,''), nama_teknisi), 
							active_yn = COALESCE(NULLIF(?,''), active_yn)
						WHERE id_teknisi = ?`

	deleteTeknisi  = "DeleteTeknisi"
	qDeleteTeknisi = "DELETE FROM teknisi where id_teknisi = ?"
)

// Add queries to key value order to be initialized as prepared statements
var (
	readStmt = []statement{
		{getAllSpareparts, qGetAllSpareparts},
		{getAllTeknisi, qGetAllTeknisi},
	}
	insertStmt = []statement{
		{createSparepart, qCreateSparepart},
		{createTeknisi, qCreateTeknisi},
	}
	updateStmt = []statement{
		{updateSparepart, qUpdateSparepart},
		{updateTeknisi, qUpdateTeknisi},
	}
	deleteStmt = []statement{
		{deleteSparepart, qDeleteSparepart},
		{deleteTeknisi, qDeleteTeknisi},
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
