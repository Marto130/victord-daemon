package service

import (
	"context"
	"victord/daemon/internal/dto"
	"victord/daemon/internal/index/models"
)

type IndexService interface {
	CreateIndex(ctx context.Context, idx *dto.CreateIndexRequest, name string) (*models.IndexResource, error)
	DestroyIndex(ctx context.Context, idx *dto.DestroyIndexRequest, name string) (*models.IndexResource, error)	
}
