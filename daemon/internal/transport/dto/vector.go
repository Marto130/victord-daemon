package dto

import (
	entity "victord/daemon/internal/entity/vector"
)

type InsertVectorRequest struct {
	ID     uint64    `json:"id"`
	Vector []float32 `json:"vector"`
}

type InsertVectorResponse struct {
	Status  string                    `json:"status"`
	Message string                    `json:"message,omitempty"`
	Results entity.InsertVectorResult `json:"results"`
}

type DeleteVectorResponse struct {
	Status  string                    `json:"status"`
	Message string                    `json:"message,omitempty"`
	Results entity.DeleteVectorResult `json:"results"`
}

type SearchVectorRequest struct {
	TopK   int       `json:"top_k"`
	Vector []float32 `json:"vector"`
}
type SearchVectorResponse struct {
	Status  string                    `json:"status"`
	Message string                    `json:"message,omitempty"`
	Results entity.SearchVectorResult `json:"results"`
}
