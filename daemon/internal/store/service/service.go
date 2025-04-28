package service

import (
	"sync"
	"victord/daemon/internal/index/models"
)

type memoryStore struct {
	indexMutex sync.RWMutex
	indexStore map[string]*models.IndexResource
}

func NewIndexStore() IndexStore {
	return &memoryStore{
		indexStore: make(map[string]*models.IndexResource),
	}
}

func (m *memoryStore) GetIndex(indexName string) (*models.IndexResource, bool) {
	m.indexMutex.RLock()
	defer m.indexMutex.RUnlock()

	index, exists := m.indexStore[indexName]
	return index, exists
}

func (m *memoryStore) GetIndexDims(indexName string) (uint16, bool) {
	m.indexMutex.RLock()
	defer m.indexMutex.RUnlock()

	index, exists := m.indexStore[indexName]
	if !exists {
		return 0, false
	}
	return index.Dims, true
}

func (m *memoryStore) StoreIndex(indexResource *models.IndexResource) {
	m.indexMutex.Lock()
	defer m.indexMutex.Unlock()

	m.indexStore[indexResource.IndexName] = indexResource
}
