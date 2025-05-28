package stock

import (
	"fmt"
	"stock/internal/entity/stock"

	"github.com/xuri/excelize/v2"
	"golang.org/x/net/context"
)

func ExportGroupedSparepartToExcel(data map[string][]stock.ReportData) ([]byte, error) {
	f := excelize.NewFile()
	sheet := "HistorySparepart"
	f.NewSheet(sheet)
	f.DeleteSheet("Sheet1")

	// Header
	headers := []string{
		"Nama Teknisi", "ID History", "ID Machine", "Counter", "Nama Sparepart",
		"Tanggal", "Quantity", "Satuan", "Total", "Nama Customer",
	}
	for i, h := range headers {
		cell := fmt.Sprintf("%s1", string('A'+i))
		f.SetCellValue(sheet, cell, h)
	}

	// Rows
	row := 2
	for teknisi, records := range data {
		for _, r := range records {
			formattedDate := r.UpdatedAt.Format("02 Jan 2006")

			f.SetCellValue(sheet, fmt.Sprintf("A%d", row), teknisi)
			f.SetCellValue(sheet, fmt.Sprintf("B%d", row), r.IDHistory)
			f.SetCellValue(sheet, fmt.Sprintf("C%d", row), r.IDMachine)
			f.SetCellValue(sheet, fmt.Sprintf("D%d", row), r.Counter)
			f.SetCellValue(sheet, fmt.Sprintf("E%d", row), r.NamaSparepart)
			f.SetCellValue(sheet, fmt.Sprintf("F%d", row), formattedDate)
			f.SetCellValue(sheet, fmt.Sprintf("G%d", row), r.Quantity)
			f.SetCellValue(sheet, fmt.Sprintf("H%d", row), r.AverageCost)
			f.SetCellValue(sheet, fmt.Sprintf("I%d", row), r.AverageCost*float64(r.Quantity))
			f.SetCellValue(sheet, fmt.Sprintf("J%d", row), r.NamaCustomer)
			row++
		}
	}

	// Write to buffer
	buffer, err := f.WriteToBuffer()
	if err != nil {
		return nil, fmt.Errorf("gagal menulis Excel ke buffer: %w", err)
	}
	return buffer.Bytes(), nil
}

func (s Service) ExportExcel(ctx context.Context, bulan, tahun int) ([]byte, string, error) {
	// 1. Ambil data dari database
	data, err := s.data.GetSparepartHistoryByMonth(ctx, bulan, tahun)
	if err != nil {
		return nil, "", fmt.Errorf("gagal ambil data: %w", err)
	}

	// 2. Kelompokkan berdasarkan nama teknisi
	grouped := make(map[string][]stock.ReportData)
	for _, item := range data {
		grouped[item.NamaTeknisi] = append(grouped[item.NamaTeknisi], item)
	}

	// 3. Tentukan nama file
	fileName := fmt.Sprintf("history_sparepart-%d-%02d.xlsx", tahun, bulan)

	// 4. Ekspor ke Excel
	fileBytes, err := ExportGroupedSparepartToExcel(grouped)
	if err != nil {
		return nil, "", err
	}

	return fileBytes, fileName, nil
}
