// backend/internal/router/router.go
package router

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourname/csvxlchart/backend/internal/config"
	"github.com/yourname/csvxlchart/backend/internal/handlers"
	"github.com/yourname/csvxlchart/backend/internal/middleware"
	"github.com/yourname/csvxlchart/backend/internal/services"
	"github.com/yourname/csvxlchart/backend/internal/storage"
)

func New(cfg config.Config) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery(), middleware.Logging(), middleware.CORS(cfg.AllowOrigin))

	// === pilih salah satu store ===

	// 1. In-Memory (cepat, cocok dev/testing)
	// store := storage.NewMemStore()

	// 2. Disk (persisten, cocok production kecil/medium)
	store, err := storage.NewDiskStore("./tmpdata")
	if err != nil {
		log.Fatalf("cannot init disk store: %v", err)
	}

	parseSvc := services.NewParseService()
	chartSvc := services.NewChartService()
	llmSvc := services.NewLLMService(cfg.OpenAIKey)

	uh := handlers.NewUploadHandler(parseSvc, store)
	ch := handlers.NewChartHandler(store, chartSvc, llmSvc)

	r.GET("/healthz", func(c *gin.Context) { c.String(http.StatusOK, "ok") })
	r.POST("/upload", uh.HandleUpload)
	r.POST("/chart", ch.HandleChart)

	return r
}
