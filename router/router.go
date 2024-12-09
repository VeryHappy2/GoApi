package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
)

type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

func NewRouter(routes []Route) *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.RealIP)

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))

	for _, route := range routes {
		router.Method(route.Method, route.Path, route.HandlerFunc)
		log.Printf("Route added: %#v\n", route)
	}

	return router
}
