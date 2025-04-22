package models

import "victord/daemon/platform/victor"

type IndexResource struct {
	IndexID   string `json:"index_id"`
	IndexType int    `json:"index_type"`
	Method    int    `json:"method"`
	Dims      uint16 `json:"dims"`
	IndexName string `json:"index_name"`
	VIndex    *victor.Index
}
