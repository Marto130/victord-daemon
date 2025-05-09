package models

import (
	wrapper "victord/daemon/internal/nativeops/cimpl"
)

type IndexResource struct {
	IndexID   string `json:"index_id"`
	IndexType int    `json:"index_type"`
	Method    int    `json:"method"`
	Dims      uint16 `json:"dims"`
	IndexName string `json:"index_name"`
	VIndex    *wrapper.VIndex
}
