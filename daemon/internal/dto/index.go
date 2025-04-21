package dto

import (
	entity "victord/daemon/internal/entity/index"
)

type CreateIndexRequest struct {
	IndexType int    `json:"index_type"`
	Method    int    `json:"method"`
	Dims      uint16 `json:"dims"`
}

type CreateIndexResponse struct {
	Status  string                   `json:"status"`
	Message string                   `json:"message,omitempty"`
	Results entity.CreateIndexResult `json:"results"`
}
