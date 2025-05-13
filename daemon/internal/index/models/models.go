package models

import (
	"victord/daemon/internal/index/factory"
	"victord/daemon/internal/nativeops/cimpl"
)

type IndexResource struct {
	IndexID   string             `json:"index_id"`
	IndexType factory.IndexType  `json:"index_type"`
	Method    factory.MethodType `json:"method"`
	Dims      uint16             `json:"dims"`
	IndexName string             `json:"index_name"`
	VIndex    *cimpl.VIndex
}
