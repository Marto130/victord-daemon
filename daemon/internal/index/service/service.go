package service

import (
	"context"
	"errors"
	"victord/daemon/internal/dto"
	"victord/daemon/internal/entity/index"
	"victord/daemon/internal/index/factory"
	"victord/daemon/internal/index/models"
	"victord/daemon/internal/nativeops"
	"victord/daemon/internal/store/service"

	"github.com/google/uuid"
)

type indexService struct {
	store    service.IndexStore
	index    factory.IndexFactory
	indexOps nativeops.IndexOps
}

func NewIndexService(store service.IndexStore, index factory.IndexFactory, indexOps nativeops.IndexOps) IndexService {
	return &indexService{
		store:    store,
		index:    index,
		indexOps: indexOps,
	}
}

func (i *indexService) CreateIndex(ctx context.Context, idx *dto.CreateIndexRequest, name string) (*models.IndexResource, error) {
	genericIndex, err := i.index.CreateIndexType(idx)
	if err != nil {
		return nil, err
	}

	vindex, err := i.indexOps.AllocIndex(genericIndex)
	if err != nil {
		return nil, err
	}

	indexID := uuid.New().String()

	indexResource := &models.IndexResource{
		IndexType: factory.IndexType(idx.IndexType),
		Method:    factory.MethodType(idx.Method),
		Dims:      idx.Dims,
		VIndex:    vindex,
		IndexName: name,
		IndexID:   indexID,
	}

	i.store.StoreIndex(indexResource)

	return indexResource, err
}

func (i *indexService) DestroyIndex(ctx context.Context, name string) (*index.DestroyIndexResult, error) {

	indexResource, exists := i.store.GetIndex(name)
	if !exists {
		return nil, errors.New("Index not found")
	}

	//TODO: Here we need to retrieve a message from the binding if the index doesn't exists.
	indexResource.VIndex.DestroyIndex()

	destroyResult := index.DestroyIndexResult{
		ID:        indexResource.IndexID,
		IndexName: indexResource.IndexName,
	}

	return &destroyResult, nil
}
