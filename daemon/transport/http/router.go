package api

import (
	"net/http"

	"victord/daemon/internal/index/factory"
	iService "victord/daemon/internal/index/service"
	"victord/daemon/internal/nativeops/cimpl"
	"victord/daemon/internal/store/service"
	vService "victord/daemon/internal/vector/service"
	"victord/daemon/transport/http/handlers"
	"victord/daemon/transport/http/routes"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, h *handlers.Handler) {
	r.HandleFunc(routes.IndexPath, h.CreateIndexHandler).Methods(http.MethodPost)
	r.HandleFunc(routes.IndexPath, h.DestroyIndexHandler).Methods(http.MethodDelete)
	r.HandleFunc(routes.InsertVectorPath, h.InsertVectorHandler).Methods(http.MethodPost)
	r.HandleFunc(routes.DeleteVectorPath, h.DeleteVectorHandler).Methods(http.MethodDelete)
	r.HandleFunc(routes.SearchVectorPath, h.SearchVectorHandler).Methods(http.MethodGet)
}

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	indexStore := service.NewIndexStore()
	indexFactory := factory.NewIndexFactory()
	indexOps := cimpl.NewIndexConstructor()

	handler := &handlers.Handler{
		IndexService:  iService.NewIndexService(indexStore, indexFactory, indexOps),
		VectorService: vService.NewVectorService(indexStore),
	}
	RegisterRoutes(router, handler)
	return router
}
