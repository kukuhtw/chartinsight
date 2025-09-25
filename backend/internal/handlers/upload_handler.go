// backend/internal/handlers/upload_handler.go
package handlers

import (
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yourname/csvxlchart/backend/internal/models"
	"github.com/yourname/csvxlchart/backend/internal/services"
	"github.com/yourname/csvxlchart/backend/internal/storage"
)

type UploadHandler struct {
	parse services.ParseService
	store storage.Store
}

func NewUploadHandler(p services.ParseService, s storage.Store) *UploadHandler {
	return &UploadHandler{parse: p, store: s}
}

func (h *UploadHandler) HandleUpload(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "file is required"})
		return
	}
	f, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "cannot open file"})
		return
	}
	defer f.Close()

	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	var ds *models.Dataset
	switch ext {
	case ".csv":
		ds, err = h.parse.ParseCSV(f)
	case ".xls", ".xlsx":
		ds, err = h.parse.ParseXLSX(toReader(f))
	default:
		err = errors.New("unsupported format: use CSV/XLS/XLSX")
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := uuid.NewString()
	h.store.Save(id, ds)

	c.JSON(http.StatusOK, gin.H{
		"uploadID": id,
		"columns":  ds.Cols,
		"rows":     len(ds.Rows),
	})
}

// convert multipart.File to io.Reader (we'll read all for excelize)
func toReader(mf multipart.File) io.Reader {
	// excelize.OpenReader needs an io.Reader. The multipart.File already satisfies it.
	// Some drivers require re-open; here we just pass it through.
	return mf
}
