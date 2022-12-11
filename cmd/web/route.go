package main

import (
	"github.com/KyawKyawThar/gowebtest/pkg/config"
	"github.com/KyawKyawThar/gowebtest/pkg/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

// Handler responds to an HTTP request.
func routes(config *config.AppConfig) http.Handler {

	mux := chi.NewRouter()
	mux.Use(Nosurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
