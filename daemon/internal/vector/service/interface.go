package service

import (
	"victord/daemon/internal/dto"
	vectorEntity "victord/daemon/internal/entity/vector"
)

type VectorService interface {
	InsertVector(*dto.InsertVectorRequest, string) (*uint64, error)
	DeleteVector(uint64, string) error
	SearchVector([]float32, string, int) (*vectorEntity.SearchVectorResult, error)
}
