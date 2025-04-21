package api

import (
	"net/http"
	iService "victord/daemon/internal/index/service"
	vService "victord/daemon/internal/vector/service"
	"victord/daemon/transport/http/handlers"
	"victord/daemon/transport/http/routes"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, h *handlers.Handler) {
	r.HandleFunc(routes.CreateIndexPath, h.CreateIndexHandler).Methods(http.MethodPost)
	r.HandleFunc(routes.InsertVectorPath, h.InsertVectorHandler).Methods(http.MethodPost)
	r.HandleFunc(routes.DeleteVectorPath, h.DeleteVectorHandler).Methods(http.MethodDelete)
	r.HandleFunc(routes.SearchVectorPath, h.SearchVectorHandler).Methods(http.MethodGet)
}

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	handler := &handlers.Handler{
		IndexService:  iService.NewIndexService(),
		VectorService: vService.NewVectorService(),
	}
	RegisterRoutes(router, handler)
	return router
}
