package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range ApplicationRoutes {

		var handler http.Handler = route.Handler;

		router.
			Methods(route.Method).
			Path(route.Path).
			Name(route.Name).
			Handler(handler)

	}

	return router
}