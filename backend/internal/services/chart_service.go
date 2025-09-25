// backend/internal/services/chart_service.go
package services

import (
	"math"
	"sort"
	"strconv"
	"strings"

	"github.com/yourname/csvxlchart/backend/internal/models"
	"github.com/yourname/csvxlchart/backend/internal/utils"
)

type ChartService interface {
	ExtractXY(ds *models.Dataset, colX, colY string) (xs []string, ys []float64)
	QuickStats(vs []float64) models.Stats
	// baru:
	GroupAggregate(ds *models.Dataset, colX, colY, groupBy, agg string) (xLabels []string, series []models.Series)
}

type chartService struct{}

func NewChartService() ChartService { return &chartService{} }

func (s *chartService) ExtractXY(ds *models.Dataset, colX, colY string) ([]string, []float64) {
	ix := ds.IndexOf(colX)
	iy := ds.IndexOf(colY)
	if ix < 0 || iy < 0 {
		return nil, nil
	}
	xs := make([]string, 0, len(ds.Rows))
	ys := make([]float64, 0, len(ds.Rows))
	for _, r := range ds.Rows {
		if ix >= len(r) || iy >= len(r) {
			continue
		}
		xs = append(xs, r[ix])
		val := strings.TrimSpace(r[iy])
		if v, err := strconv.ParseFloat(val, 64); err == nil {
			ys = append(ys, v)
		}
	}
	return xs, ys
}

func (s *chartService) QuickStats(vs []float64) models.Stats {
	if len(vs) == 0 {
		return models.Stats{}
	}
	min, max, sum := vs[0], vs[0], 0.0
	for _, v := range vs {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
		sum += v
	}
	return models.Stats{
		N:    len(vs),
		Min:  min,
		Max:  max,
		Mean: sum / float64(len(vs)),
		Std:  utils.Std(vs),
	}
}

// GroupAggregate: untuk setiap nilai X, hitung agregasi(Y) per kelompok (GroupBy)
func (s *chartService) GroupAggregate(ds *models.Dataset, colX, colY, groupBy, agg string) ([]string, []models.Series) {
	ix := ds.IndexOf(colX)
	iy := ds.IndexOf(colY)
	ig := ds.IndexOf(groupBy)
	if ix < 0 || iy < 0 || ig < 0 {
		return nil, nil
	}

	// Kumpulkan semua kategori X & semua nama grup
	xSet := map[string]struct{}{}
	gSet := map[string]struct{}{}
	type key struct{ x, g string }
	// pegang nilai Y per (X, Group)
	vals := map[key][]float64{}

	for _, r := range ds.Rows {
		if ix >= len(r) || iy >= len(r) || ig >= len(r) {
			continue
		}
		x := strings.TrimSpace(r[ix])
		g := strings.TrimSpace(r[ig])
		yStr := strings.TrimSpace(r[iy])
		v, err := strconv.ParseFloat(yStr, 64)
		if err != nil || math.IsNaN(v) || math.IsInf(v, 0) {
			continue
		}
		xSet[x] = struct{}{}
		gSet[g] = struct{}{}
		k := key{x: x, g: g}
		vals[k] = append(vals[k], v)
	}

	// Urutkan label X agar konsisten
	xLabels := make([]string, 0, len(xSet))
	for x := range xSet {
		xLabels = append(xLabels, x)
	}
	sort.Strings(xLabels)

	// Pastikan urutan seri per nama grup konsisten juga
	groups := make([]string, 0, len(gSet))
	for g := range gSet {
		groups = append(groups, g)
	}
	sort.Strings(groups)

	agg = strings.ToLower(strings.TrimSpace(agg))
	if agg == "" {
		agg = "avg"
	}

	series := make([]models.Series, 0, len(groups))
	for _, g := range groups {
		data := make([]float64, len(xLabels))
		for i, x := range xLabels {
			ys := vals[key{x: x, g: g}]
			data[i] = aggregate(ys, agg)
		}
		series = append(series, models.Series{Name: g, Data: data})
	}
	return xLabels, series
}

func aggregate(vs []float64, agg string) float64 {
	if len(vs) == 0 {
		return 0
	}
	switch agg {
	case "sum":
		s := 0.0
		for _, v := range vs {
			s += v
		}
		return s
	case "min":
		m := vs[0]
		for _, v := range vs {
			if v < m {
				m = v
			}
		}
		return m
	case "max":
		m := vs[0]
		for _, v := range vs {
			if v > m {
				m = v
			}
		}
		return m
	default: // "avg"
		s := 0.0
		for _, v := range vs {
			s += v
		}
		return s / float64(len(vs))
	}
}
