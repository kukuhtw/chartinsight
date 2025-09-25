// backend/internal/storage/mem_store.go
package storage

import (
	"sync"

	"github.com/yourname/csvxlchart/backend/internal/models"
)

// Store adalah kontrak minimal yang dipakai handler:
// - Save(id, ds)
// - Get(id) -> (*Dataset, bool)
type Store interface {
	Save(id string, ds *models.Dataset)
	Get(id string) (*models.Dataset, bool)
	// Opsional (tidak dipakai handler saat ini)
	Delete(id string) bool
}

type memStore struct {
	mu sync.RWMutex
	db map[string]*models.Dataset
}

// NewMemStore membuat store sederhana di memori.
func NewMemStore() Store {
	return &memStore{db: make(map[string]*models.Dataset)}
}

func (m *memStore) Save(id string, ds *models.Dataset) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.db[id] = ds
}

func (m *memStore) Get(id string) (*models.Dataset, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.db[id]
	return v, ok
}

func (m *memStore) Delete(id string) bool {
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, ok := m.db[id]; !ok {
		return false
	}
	delete(m.db, id)
	return true
}
