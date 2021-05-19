package routes

import (
	"me-english/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Uri          string
	Method       string
	Handler      func(http.ResponseWriter, *http.Request)
	AuthRequired bool
	CheckHeaders bool
}

func Load() []Route {
	routes := productRoutes
	routes = append(routes, vocabularyRoutes...)
	routes = append(routes, telegramRoutes...)
	return routes
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.Uri, route.Handler).Methods(route.Method)
	}
	return r
}

func SetupRoutesWithMiddlewares(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		if route.CheckHeaders {
			if route.AuthRequired {
				r.HandleFunc(route.Uri,
					middlewares.SetMiddlewareLogger(
						middlewares.SetMiddlewareJSON(
							middlewares.SetMiddlewareHeader(
								middlewares.SetMiddlewareVerifyToken(route.Handler))))).Methods(route.Method)
			} else {
				r.HandleFunc(route.Uri,
					middlewares.SetMiddlewareLogger(
						middlewares.SetMiddlewareJSON(
							middlewares.SetMiddlewareHeader(route.Handler)))).Methods(route.Method)
			}
		} else {
			r.HandleFunc(route.Uri,
				middlewares.SetMiddlewareLogger(
					middlewares.SetMiddlewareJSON(route.Handler))).Methods(route.Method)
		}
	}
	return r
}
