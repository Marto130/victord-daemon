package service

import (
	"context"
	"victord/daemon/internal/dto"
	"victord/daemon/internal/entity/index"
	"victord/daemon/internal/index/models"
)

type IndexService interface {
	CreateIndex(ctx context.Context, idx *dto.CreateIndexRequest, name string) (*models.IndexResource, error)
	DestroyIndex(ctx context.Context, name string) (*index.DestroyIndexResult, error)
}
