package factory

import (
	"fmt"
	"victord/daemon/internal/dto"
)

type IndexFactory interface {
	CreateIndex(*dto.CreateIndexRequest) (GenericIndex, error)
}

type defaultFactory struct{}

func NewIndexFactory() IndexFactory {
	return &defaultFactory{}
}

func (d *defaultFactory) CreateIndex(request *dto.CreateIndexRequest) (GenericIndex, error) {
	iType := IndexType(request.IndexType)
	switch iType {
	case FlatIndexType:
		return NewFlatIndex(iType, MethodType(request.Method), request.Dims), nil
	case HNSWIndexType:
		return NewHnswIndex(iType, MethodType(request.Method), request.Dims, request.Options), nil
	default:
		return nil, fmt.Errorf("unsupported index type")
	}
}
