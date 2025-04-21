package service

import (
	"victord/daemon/internal/dto"
	vectorEntity "victord/daemon/internal/entity/vector"
)

type VectorService interface {
	InsertVector(vectorData *dto.InsertVectorRequest, idxName string) (*uint64, error)
	DeleteVector(vectorId uint64, idxName string) (*uint64, error)
	SearchVector(vector []*float32, idxName string, topK int) (*vectorEntity.SearchVectorResult, error)
}
