package api

import (
	"net/http"
	"victord/daemon/internal/http/handlers"
	"victord/daemon/internal/http/routes"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	r.HandleFunc(routes.CreateIndexPath, handlers.CreateIndexHandler()).Methods(http.MethodPost)
	r.HandleFunc(routes.InsertVectorPath, handlers.InsertVectorHandler).Methods(http.MethodPost)
	r.HandleFunc(routes.DeleteVectorPath, handlers.DeleteVectorHandler).Methods(http.MethodDelete)
	r.HandleFunc(routes.SearchVectorPath, handlers.SearchVectorHandler).Methods(http.MethodGet)
}

func SetupRouter() *mux.Router {
	router := mux.NewRouter()
	RegisterRoutes(router)
	return router
}
