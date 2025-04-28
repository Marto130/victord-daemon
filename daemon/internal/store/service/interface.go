package service

import "victord/daemon/internal/index/models"

type IndexStore interface {
	GetIndex(string) (*models.IndexResource, bool)
	GetIndexDims(indexName string) (uint16, bool)
	StoreIndex(indexResource *models.IndexResource)
}
