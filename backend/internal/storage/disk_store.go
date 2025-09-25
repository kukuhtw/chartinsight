// backend/internal/storage/disk_store.go
package storage

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/yourname/csvxlchart/backend/internal/models"
)

// diskStore menyimpan Dataset sebagai file JSON di folder tertentu.
// Index ID -> path disimpan di memori agar akses cepat.
type diskStore struct {
	dir string
	mu  sync.RWMutex
	idx map[string]string // id -> filepath
}

// NewDiskStore memastikan direktori ada dan siap dipakai.
// Contoh: NewDiskStore("./tmpdata")
func NewDiskStore(dir string) (Store, error) {
	if dir == "" {
		return nil, errors.New("disk store dir is empty")
	}
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return nil, err
	}
	return &diskStore{
		dir: dir,
		idx: make(map[string]string),
	}, nil
}

func (d *diskStore) Save(id string, ds *models.Dataset) {
	// Serialize ke JSON, tulis ke file {dir}/{id}.json
	path := filepath.Join(d.dir, id+".json")

	// Tambahkan metadata minimal (optional) via wrapper
	type fileFormat struct {
		ID        string         `json:"id"`
		SavedAt   time.Time      `json:"savedAt"`
		Dataset   *models.Dataset`json:"dataset"`
		FileVer   int            `json:"fileVer"`
	}
	ff := fileFormat{
		ID:      id,
		SavedAt: time.Now().UTC(),
		Dataset: ds,
		FileVer: 1,
	}
	b, _ := json.MarshalIndent(ff, "", "  ")

	// Tulis atomically: tulis ke temp lalu rename
	tmp := path + ".tmp"
	_ = os.WriteFile(tmp, b, 0o644)
	_ = os.Rename(tmp, path)

	d.mu.Lock()
	d.idx[id] = path
	d.mu.Unlock()
}

func (d *diskStore) Get(id string) (*models.Dataset, bool) {
	d.mu.RLock()
	path, ok := d.idx[id]
	d.mu.RUnlock()
	if !ok {
		// Jika belum ada di index (misal server restart), coba cari langsung di disk
		path = filepath.Join(d.dir, id+".json")
		if _, err := os.Stat(path); err != nil {
			return nil, false
		}
		// Cache di index
		d.mu.Lock()
		d.idx[id] = path
		d.mu.Unlock()
	}

	b, err := os.ReadFile(path)
	if err != nil {
		return nil, false
	}
	var out struct {
		Dataset *models.Dataset `json:"dataset"`
	}
	if err := json.Unmarshal(b, &out); err != nil || out.Dataset == nil {
		return nil, false
	}
	return out.Dataset, true
}

func (d *diskStore) Delete(id string) bool {
	d.mu.Lock()
	path, ok := d.idx[id]
	if ok {
		delete(d.idx, id)
	}
	d.mu.Unlock()
	if !ok {
		// tetap coba hapus file dengan pola {id}.json
		path = filepath.Join(d.dir, id+".json")
	}
	_ = os.Remove(path)
	return true
}
