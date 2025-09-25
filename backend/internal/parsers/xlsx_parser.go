// backend/internal/parsers/xlsx_parser.go
package parsers

import (
	"io"

	"github.com/xuri/excelize/v2"
	"github.com/yourname/csvxlchart/backend/internal/models"
)

type XLSXParser interface {
	Read(r io.Reader) (*models.Dataset, error)
}

type xlsxParser struct{}

func NewXLSXParser() XLSXParser { return &xlsxParser{} }

func (x *xlsxParser) Read(r io.Reader) (*models.Dataset, error) {
	f, err := excelize.OpenReader(r)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return &models.Dataset{}, nil
	}
	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return nil, err
	}
	return models.FromRecords(rows), nil
}
