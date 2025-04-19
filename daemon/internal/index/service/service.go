package service

import (
	binding "victord/binding"
	"victord/daemon/internal/index/models"
	"victord/daemon/internal/store/service"
	"victord/daemon/internal/transport/dto"

	"github.com/google/uuid"
)

func CreateIndex(idx *dto.CreateIndexRequest, name string) (*models.IndexResource, error) {

	index, err := binding.AllocIndex(idx.IndexType, idx.Method, idx.Dims)
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
