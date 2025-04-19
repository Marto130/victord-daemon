package handlers

import (
	"encoding/json"
	"net/http"
	"sync"

	indexEntity "victord/daemon/internal/entity/index"
	"victord/daemon/internal/index/service"
	"victord/daemon/internal/transport/dto"

	"github.com/gorilla/mux"
)

var (
	indexMutex sync.RWMutex
)

func CreateIndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var createIndexRequest dto.CreateIndexRequest

		urlParams := mux.Vars(r)
		indexNameParam := urlParams["indexName"]
		if indexNameParam == "" {
			http.Error(w, "Index name is required", http.StatusBadRequest)
			return
		}

		indexMutex.Lock()
		defer indexMutex.Unlock()

		if err := json.NewDecoder(r.Body).Decode(&createIndexRequest); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		indexResource, err := service.CreateIndex(&createIndexRequest, indexNameParam)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)

		}

		w.Header().Set("Content-Type", "application/json")

		res := dto.CreateIndexResponse{
			Status:  "success",
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
}
