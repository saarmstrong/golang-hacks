package lib

import (
	"net/http"

	"github.com/gorilla/mux"
)

var GRobot *LilRobots

func NewRouter(rb *LilRobots) *mux.Router {
	GRobot = rb
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler

		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}

	return router
}
