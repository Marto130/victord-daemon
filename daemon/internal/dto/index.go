package dto

import (
	entity "victord/daemon/internal/entity/index"
)

type CreateIndexRequest struct {
	IndexType int            `json:"index_type"`
	Method    int            `json:"method"`
	Dims      uint16         `json:"dims"`
	Options   map[string]int `json:"options"` //TODO: generic type
}

type CreateIndexResponse struct {
	Status  string                   `json:"status"`
	Message string                   `json:"message,omitempty"`
	Results entity.CreateIndexResult `json:"results"`
}

type DestroyIndexResponse struct {
	Status  string                    `json:"status"`
	Message string                    `json:"message,omitempty"`
	Results entity.DestroyIndexResult `json:"results"`
}
