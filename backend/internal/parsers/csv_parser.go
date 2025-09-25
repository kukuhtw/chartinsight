// backend/internal/parsers/csv_parser.go
package parsers

type CSVParser interface{}
func NewCSVParser() CSVParser { return &csvParser{} }
type csvParser struct{}
