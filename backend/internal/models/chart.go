// backend/internal/models/chart.go
package models

type ChartRequest struct {
	UploadID string `json:"uploadID"`
	ColX     string `json:"colX"`
	ColY     string `json:"colY"`
	GroupBy  string `json:"groupBy,omitempty"` // <- baru, opsional
	Agg      string `json:"agg,omitempty"`     // <- baru: "sum"|"avg"|"min"|"max" (default: "avg")
}

type Stats struct {
	N    int     `json:"n"`
	Min  float64 `json:"min"`
	Max  float64 `json:"max"`
	Mean float64 `json:"mean"`
	Std  float64 `json:"std"`
}

// Untuk grouped chart: tiap grup punya satu seri nilai Y di atas sumbu X
type Series struct {
	Name string    `json:"name"` // nama grup
	Data []float64 `json:"data"` // selaras dengan urutan XLabels
}

type ChartResponse struct {
	// Legacy (single series)
	X       []string  `json:"x,omitempty"`
	Y       []float64 `json:"y,omitempty"`

	// Grouped
	XLabels []string `json:"xLabels,omitempty"`
	Series  []Series `json:"series,omitempty"`

	Stats   Stats             `json:"stats"`
	Insight string            `json:"insight"`
	EChartsOption map[string]any `json:"echartsOption,omitempty"`
}
