// backend/internal/utils/stats.go
package utils

import "math"

func Std(vs []float64) float64 {
	if len(vs) == 0 {
		return 0
	}
	m := mean(vs)
	var sum float64
	for _, v := range vs {
		d := v - m
		sum += d * d
	}
	return math.Sqrt(sum / float64(len(vs)))
}

func mean(vs []float64) float64 {
	if len(vs) == 0 {
		return 0
	}
	var s float64
	for _, v := range vs {
		s += v
	}
	return s / float64(len(vs))
}
