package api

import (
	"log"
	"net/http"
	"time"

	"transport-backend/pkg/httpx"
	"transport-backend/pkg/routes"
	"transport-backend/pkg/trips"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

var router *chi.Mux

func init() {
	router = chi.NewRouter()

	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.Timeout(30 * time.Second))

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	router.Get("/health", healthHandler)
	router.Get("/routes", routes.ListRoutes)
	router.Get("/routes/{id}", routes.GetRoute)
	router.Post("/trips", trips.CreateTrip)
	router.Get("/trips/{id}", trips.GetTrip)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	httpx.JSON(w, http.StatusOK, map[string]string{
		"status":    "ok",
		"timestamp": time.Now().UTC().Format(time.RFC3339),
	})
}

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Printf("[%s] %s", r.Method, r.URL.Path)
	router.ServeHTTP(w, r)
}
