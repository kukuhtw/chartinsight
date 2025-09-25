// backend/internal/models/dataset.go
package models

type Dataset struct {
	Cols []string   `json:"cols"`
	Rows [][]string `json:"rows"`
}

func (d *Dataset) IndexOf(col string) int {
	for i, c := range d.Cols {
		if c == col {
			return i
		}
	}
	return -1
}

func FromRecords(records [][]string) *Dataset {
	if len(records) == 0 {
		return &Dataset{}
	}
	if len(records) == 1 {
		return &Dataset{Cols: records[0], Rows: [][]string{}}
	}
	return &Dataset{
		Cols: records[0],
		Rows: records[1:],
	}
}
