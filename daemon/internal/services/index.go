package services

import (
	binding "victord/binding"
	"victord/daemon/pkg/models"

	"github.com/google/uuid"
)

func CreateIndex(idx *models.CreateIndexRequest, name string) (*models.IndexResource, error) {

	index, err := binding.AllocIndex(idx.IndexType, idx.Method, idx.Dims)
	if err != nil {
		return nil, err
	}

	indexID := uuid.New().String()

	indexResource := models.IndexResource{
		CreateIndexRequest: models.CreateIndexRequest{
			IndexType: idx.IndexType,
			Method:    idx.Method,
			Dims:      idx.Dims,
		},
		VIndex:    index,
		IndexName: name,
		IndexID:   indexID,
	}

	StoreIndex(&indexResource)

	return &indexResource, err
}
