package models

import binding "victord/binding"

type IndexResource struct {
	CreateIndexRequest
	IndexID   string `json:"index_id"`
	IndexName string `json:"index_name"`
	VIndex    *binding.Index
}
