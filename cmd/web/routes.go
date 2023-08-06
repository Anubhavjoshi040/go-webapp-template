package main

import (
	"net/http"

	"github.com/anubhavjoshi040/go-webapp-template/config"
	"github.com/anubhavjoshi040/go-webapp-template/pkg/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)


	return mux
}