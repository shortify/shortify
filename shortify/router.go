package shortify

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type route struct {
	name        string
	method      string
	pattern     string
	handlerFunc http.HandlerFunc
}

var routes []route
var shortifyPort string

func init() {
	routes = []route{
		route{"redirect", "GET", "/{token}", performRedirectHandler},
		route{"create", "POST", "/redirects", createRedirectHandler},
	}
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		handler := logger(route.handlerFunc, route.name)
		router.
			Methods(route.method).
			Path(route.pattern).
			Name(route.name).
			Handler(handler)
	}

	return router
}

func ServerPort() string {
	return fmt.Sprintf(":%s", shortifyPort)
}

func logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
