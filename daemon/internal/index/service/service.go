package service

import (
	"context"
	"victord/daemon/internal/dto"
	"victord/daemon/internal/index/models"
	"victord/daemon/internal/store/service"
	"victord/daemon/platform/victor"

	"github.com/google/uuid"
)

type indexService struct {
}

func NewIndexService() IndexService {
	return &indexService{}
}

func (i *indexService) CreateIndex(ctx context.Context, idx *dto.CreateIndexRequest, name string) (*models.IndexResource, error) {

	index, err := victor.AllocIndex(idx.IndexType, idx.Method, idx.Dims)
	if err != nil {
		return nil, err
	}

	indexID := uuid.New().String()

	indexResource := models.IndexResource{
		IndexType: idx.IndexType,
		Method:    idx.Method,
		Dims:      idx.Dims,
		VIndex:    index,
		IndexName: name,
		IndexID:   indexID,
	}

	service.StoreIndex(&indexResource)

	return &indexResource, err
}
