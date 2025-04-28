package api

import (
	"net/http"

	iService "victord/daemon/internal/index/service"
	c "victord/daemon/internal/nativeops/cimpl"
	"victord/daemon/internal/store/service"
	vService "victord/daemon/internal/vector/service"
	"victord/daemon/transport/http/handlers"

	"github.com/gorilla/mux"
)

const (
	CreateIndexPath  = "/api/index/{indexName}"
	InsertVectorPath = "/api/vector/{indexName}"
	DeleteVectorPath = "/api/vector/{indexName}/{vectorID}"
	SearchVectorPath = "/api/vector/{indexName}/search"
)

func RegisterRoutes(r *mux.Router, h *handlers.Handler) {
	r.HandleFunc(CreateIndexPath, h.CreateIndexHandler).Methods(http.MethodPost)
	r.HandleFunc(InsertVectorPath, h.InsertVectorHandler).Methods(http.MethodPost)
	r.HandleFunc(DeleteVectorPath, h.DeleteVectorHandler).Methods(http.MethodDelete)
	r.HandleFunc(SearchVectorPath, h.SearchVectorHandler).Methods(http.MethodGet)
}

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	indexStore := service.NewIndexStore()
	indexOps := c.NewIndex()
	handler := &handlers.Handler{
		IndexService:  iService.NewIndexService(indexStore, indexOps),
		VectorService: vService.NewVectorService(indexStore),
	}
	RegisterRoutes(router, handler)
	return router
}
