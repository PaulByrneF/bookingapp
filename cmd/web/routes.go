package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/paulbyrne/bookingapp/pkg/config"
	"github.com/paulbyrne/bookingapp/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {

	mux := chi.NewRouter()

	//prints stack trace when panic & recover application
	mux.Use(middleware.Recoverer)

	// intercept and print request url
	mux.Use(WriteToConsole)

	// csrf protection
	mux.Use(NoSurf)

	// access session
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// fileserver for serving static files
	fileserver := http.FileServer(http.Dir("../../static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileserver))

	return mux
}
