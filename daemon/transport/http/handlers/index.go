package handlers

import (
	"encoding/json"
	"net/http"
	"victord/daemon/internal/dto"
	indexEntity "victord/daemon/internal/entity/index"

	"github.com/gorilla/mux"
)

func (h *Handler) CreateIndexHandler(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	indexNameParam := urlParams["indexName"]

	var createIndexRequest dto.CreateIndexRequest
	if err := json.NewDecoder(r.Body).Decode(&createIndexRequest); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	indexResource, err := h.IndexService.CreateIndex(r.Context(), &createIndexRequest, indexNameParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := &dto.CreateIndexResponse{
		Status:  "Success",
		Message: "Index created successfully",
		Results: indexEntity.CreateIndexResult{
			IndexName: indexResource.IndexName,
			ID:        indexResource.IndexID,
			Dims:      indexResource.Dims,
			IndexType: int(indexResource.IndexType),
			Method:    int(indexResource.Method),
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) DestroyIndexHandler(w http.ResponseWriter, r *http.Request) {

	urlParams := mux.Vars(r)
	indexNameParam := urlParams["indexName"]

	destroyResult, err := h.IndexService.DestroyIndex(r.Context(), indexNameParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := dto.DestroyIndexResponse{
		Status:  "Success",
		Message: "Index destroyed successfully",
		Results: indexEntity.DestroyIndexResult{
			ID:        destroyResult.ID,
			IndexName: destroyResult.IndexName,
		},
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
