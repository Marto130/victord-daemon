package service

import (
	"context"
	"victord/daemon/internal/dto"
	"victord/daemon/internal/index/models"
	"victord/daemon/internal/nativeops"
	"victord/daemon/internal/store/service"

	"github.com/google/uuid"
)

type indexService struct {
	store    service.IndexStore
	indexOps nativeops.IndexOps
}

func NewIndexService(store service.IndexStore, indexOps nativeops.IndexOps) IndexService {
	return &indexService{
		store:    store,
		indexOps: indexOps,
	}
}

func (i *indexService) CreateIndex(ctx context.Context, idx *dto.CreateIndexRequest, name string) (*models.IndexResource, error) {
	index, err := i.indexOps.AllocIndex(idx.IndexType, idx.Method, idx.Dims)
	if err != nil {
		return nil, err
	}

	indexID := uuid.New().String()

	indexResource := &models.IndexResource{
		IndexType: idx.IndexType,
		Method:    idx.Method,
		Dims:      idx.Dims,
		VIndex:    index,
		IndexName: name,
		IndexID:   indexID,
	}

	i.store.StoreIndex(indexResource)

	return indexResource, err
}
