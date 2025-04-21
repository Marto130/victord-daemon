package api

import (
	"net/http"
	"victord/daemon/internal/index/service"
	"victord/daemon/transport/http/handlers"
	"victord/daemon/transport/http/routes"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, h *handlers.Handler) {
	r.HandleFunc(routes.CreateIndexPath, h.CreateIndexHandler).Methods(http.MethodPost)
	r.HandleFunc(routes.InsertVectorPath, handlers.InsertVectorHandler).Methods(http.MethodPost)
	r.HandleFunc(routes.DeleteVectorPath, handlers.DeleteVectorHandler).Methods(http.MethodDelete)
	r.HandleFunc(routes.SearchVectorPath, handlers.SearchVectorHandler).Methods(http.MethodGet)
}

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	indexService := service.NewIndexService()
	handler := &handlers.Handler{
		IndexService: indexService,
	}
	RegisterRoutes(router, handler)
	return router
}
