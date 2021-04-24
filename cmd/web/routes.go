package main

import (
	"net/http"

	"github.com/go-chi/chi/v5/middleware"

	"github.com/danarcheronline/wizard-bookings/pkg/handlers"

	"github.com/go-chi/chi/v5"

	"github.com/danarcheronline/wizard-bookings/pkg/config"
)

func routes(*config.AppConfig) http.Handler {
	// mux := pat.New()
	// mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	// mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	return mux
}
