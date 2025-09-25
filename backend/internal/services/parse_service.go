// backend/internal/services/parse_service.go
package services

import (
	"bytes"
	"encoding/csv"
	"io"

	"github.com/yourname/csvxlchart/backend/internal/models"
	"github.com/yourname/csvxlchart/backend/internal/parsers"
)

type ParseService interface {
	ParseCSV(r io.Reader) (*models.Dataset, error)
	ParseXLSX(r io.Reader) (*models.Dataset, error)
}

type parseService struct {
	csvP  parsers.CSVParser
	xlsxP parsers.XLSXParser
}

func NewParseService() ParseService {
	return &parseService{
		csvP:  parsers.NewCSVParser(),
		xlsxP: parsers.NewXLSXParser(),
	}
}

func (s *parseService) ParseCSV(r io.Reader) (*models.Dataset, error) {
	cr := csv.NewReader(r)
	cr.FieldsPerRecord = -1
	records, err := cr.ReadAll()
	if err != nil {
		return nil, err
	}
	return models.FromRecords(records), nil
}

func (s *parseService) ParseXLSX(r io.Reader) (*models.Dataset, error) {
	// read all for excelize reader
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, r); err != nil {
		return nil, err
	}
	return s.xlsxP.Read(bytes.NewReader(buf.Bytes()))
}
