package handlers

import (
	"encoding/json"
	"net/http"	
	"victord/daemon/internal/dto"
	indexEntity "victord/daemon/internal/entity/index"
	"github.com/gorilla/mux"
)

func (h *Handler) CreateIndexHandler(w http.ResponseWriter, r *http.Request) {

	var createIndexRequest dto.CreateIndexRequest

	urlParams := mux.Vars(r)
	indexNameParam := urlParams["indexName"]
	if indexNameParam == "" {
		http.Error(w, "Index name is required", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&createIndexRequest); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	indexResource, err := h.IndexService.CreateIndex(r.Context(), &createIndexRequest, indexNameParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")

	res := dto.CreateIndexResponse{
		Status:  "Success",
		Message: "Index created successfully",
		Results: indexEntity.CreateIndexResult{
			IndexName: indexResource.IndexName,
			ID:        indexResource.IndexID,
			Dims:      indexResource.Dims,
			IndexType: indexResource.IndexType,
			Method:    indexResource.Method,
		},
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func (h *Handler) DestroyIndexHandler(w http.ResponseWriter, r *http.Request) {
	
	urlParams := mux.Vars(r);
	var destroyIndexRequest dto.DestroyIndexRequest
	indexNameParam := urlParams["indexName"]
	if indexNameParam == "" {
		http.Error(w, "Index name is required", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&destroyIndexRequest); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	indexResource, err := h.IndexService.DestroyIndex(r.Context(), &destroyIndexRequest, indexNameParam)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

	}

	w.Header().Set("Content-Type", "application/json")

	res := dto.DestroyIndexResponse{
		Status:  "Success",
		Message: "Index destroyed successfully",
		Results: indexEntity.DestroyIndexResult{
			IndexName: indexResource.IndexName,
			ID:        indexResource.IndexID,
			Dims:      indexResource.Dims,
			IndexType: indexResource.IndexType,
			Method:    indexResource.Method,
		},
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
