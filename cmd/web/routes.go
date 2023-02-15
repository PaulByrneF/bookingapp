package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/paulbyrne/bookingapp/internal/config"
	"github.com/paulbyrne/bookingapp/internal/handlers"
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

	mux.Get("/about-us", handlers.Repo.AboutUs)

	mux.Get("/search-availability", handlers.Repo.SearchAvailability)
	mux.Post("/check-availability", handlers.Repo.PostCheckAvailability)
	mux.Post("/check-availability-json", handlers.Repo.AvailabilityJSON)

	mux.Get("/make-reservation", handlers.Repo.MakeReservation)

	mux.Get("/contact-us", handlers.Repo.ContactUs)

	mux.Get("/rooms/basic-suite", handlers.Repo.BasicSuite)
	mux.Get("/rooms/kings-suite", handlers.Repo.KingsSuite)

	// fileserver for serving static files
	fileserver := http.FileServer(http.Dir("../../static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileserver))

	return mux
}
