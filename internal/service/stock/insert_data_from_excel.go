package stock

import (
	"context"
	"log"
	"stock/internal/entity/stock"
	"stock/pkg/errors"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func (s Service) ImportCustomersFromExcel(ctx context.Context) error {
	filePath := `C:\Users\user\go\src\stock\internal\service\stock\Customer All PB GROUP (Wilayah).xlsx`

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][ImportCustomersFromExcel][OPEN_FILE]")
	}
	defer f.Close()

	sheetNames := f.GetSheetList()
	if len(sheetNames) == 0 {
		return errors.New("[SERVICE][ImportCustomersFromExcel][NO_SHEETS_FOUND]")
	}
	sheetName := sheetNames[6]

	rows, err := f.Rows(sheetName)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][ImportCustomersFromExcel][GET_ROWS_ITERATOR]")
	}
	defer rows.Close()

	if !rows.Next() {
		return errors.New("[SERVICE][ImportCustomersFromExcel][HEADER_NOT_FOUND_OR_EMPTY]")
	}

	rowNumber := 1

	for rows.Next() {
		rowNumber++
		row, err := rows.Columns()
		if err != nil {
			log.Printf("Warning: Gagal membaca kolom di baris %d. Error: %v. Baris dilewati.", rowNumber, err)
			continue
		}

		if len(row) < 7 || row[1] == "" || row[6] == "" {
			log.Printf("Warning: Baris %d tidak lengkap atau data penting kosong. Baris dilewati.", rowNumber)
			continue
		}

		comID, err := strconv.Atoi(row[6])
		if err != nil {
			log.Printf("Warning: Gagal konversi Company ID di baris %d (value: '%s'). Error: %v. Baris dilewati.", rowNumber, row[6], err)
			continue
		}

		customer := stock.Customer{
			NamaCustomer: row[1],
			Company:      comID,
			Alamat:       row[2],
			UpdatedBy:    "admin",
		}

		err = s.CreateCustomer(ctx, customer)
		if err != nil {
			return errors.Wrap(err, "[SERVICE][CreateCustomer]")
		}
	}

	return nil
}

func (s Service) ImportMachineFromExcel(ctx context.Context) error {
	filePath := `C:\Users\user\go\src\stock\internal\service\stock\Customer All PB GROUP (Wilayah).xlsx`

	f, err := excelize.OpenFile(filePath)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][ImportCustomersFromExcel][OPEN_FILE]")
	}
	defer f.Close()

	sheetNames := f.GetSheetList()
	if len(sheetNames) == 0 {
		return errors.New("[SERVICE][ImportCustomersFromExcel][NO_SHEETS_FOUND]")
	}
	sheetName := sheetNames[36]

	rows, err := f.Rows(sheetName)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][ImportCustomersFromExcel][GET_ROWS_ITERATOR]")
	}
	defer rows.Close()

	if !rows.Next() {
		return errors.New("[SERVICE][ImportCustomersFromExcel][HEADER_NOT_FOUND_OR_EMPTY]")
	}

	rowNumber := 1

	for rows.Next() {
		rowNumber++
		row, err := rows.Columns()
		if len(row) < 9 {
			log.Printf("Warning: Baris %d hanya memiliki %d kolom. Baris dilewati.", rowNumber, len(row))
			continue
		}
		if row[8] == "" {
			log.Printf("Warning: Baris %d kolom[8] kosong. Baris dilewati.", rowNumber)
			continue
		}

		machine := stock.Machine{
			ID:           row[5],
			Type:         row[3],
			SerialNumber: row[4],
			Customer:     row[7],
		}

		history := stock.MachineHistory{
			IDMachine:          row[5],
			IDCustomer:         row[7],
			TanggalMulaiString: row[8],
			Status:             "Aktif",
		}

		err = s.CreateMachine(ctx, machine)
		if err != nil {
			return errors.Wrap(err, "[SERVICE][CreateMachine]")
		}

		err = s.CreateMachineHistory(ctx, history)
		if err != nil {
			return errors.Wrap(err, "[SERVICE][CreateMachineHistory]")
		}
	}

	return nil
}
