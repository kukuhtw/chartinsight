// backend/internal/handlers/chart_handler.go
package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourname/csvxlchart/backend/internal/models"
	"github.com/yourname/csvxlchart/backend/internal/services"
	"github.com/yourname/csvxlchart/backend/internal/storage"
)

type ChartHandler struct {
	store    storage.Store
	chartSvc services.ChartService
	llmSvc   services.LLMService
}

func NewChartHandler(s storage.Store, c services.ChartService, l services.LLMService) *ChartHandler {
	return &ChartHandler{store: s, chartSvc: c, llmSvc: l}
}

func (h *ChartHandler) HandleChart(c *gin.Context) {
	var req models.ChartRequest
	if err := c.ShouldBindJSON(&req); err != nil || req.UploadID == "" || req.ColX == "" || req.ColY == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	ds, ok := h.store.Get(req.UploadID)
	if !ok || ds == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "upload not found"})
		return
	}

	// grouped?
	if req.GroupBy != "" {
		xLabels, series := h.chartSvc.GroupAggregate(ds, req.ColX, req.ColY, req.GroupBy, req.Agg)
		// hitung stats dari semua nilai yang terbentuk (dirata-rata semua seri)
		all := make([]float64, 0, len(xLabels)*len(series))
		for _, s := range series {
			all = append(all, s.Data...)
		}
		stats := h.chartSvc.QuickStats(all)
		insight := h.llmSvc.InsightGrouped(req.ColX, req.ColY, req.GroupBy, req.Agg, xLabels, series, stats)

		c.JSON(http.StatusOK, models.ChartResponse{
			XLabels: xLabels,
			Series:  series,
			Stats:   stats,
			Insight: insight,
		})
		return
	}

	// non-grouped (single series)
	xs, ys := h.chartSvc.ExtractXY(ds, req.ColX, req.ColY)
	stats := h.chartSvc.QuickStats(ys)
	insight := h.llmSvc.Insight(req.ColX, req.ColY, xs, ys, stats)

	c.JSON(http.StatusOK, models.ChartResponse{
		X:       xs,
		Y:       ys,
		Stats:   stats,
		Insight: insight,
	})
}
