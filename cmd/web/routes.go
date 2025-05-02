package main

import (
	"net/http"
	"github.com/ZackDiego/bookings/pkg/config"
	"github.com/ZackDiego/bookings/pkg/handlers"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func routes(app *config.AppConfig) http.Handler {
	// router := mux.NewRouter()

	// router.HandleFunc("/", handlers.Repo.Home).Methods("GET")
	// router.HandleFunc("/about", handlers.Repo.About).Methods("GET")

	router := chi.NewRouter()

	router.Use(middleware.Recoverer)
	router.Use(WriteToConsole)
	router.Use(NoSurf)
	router.Use(SessionLoad)
	
	router.Get("/", handlers.Repo.Home)
	router.Get("/about", handlers.Repo.About)
	
	return router
}

