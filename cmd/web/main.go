package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/paulbyrne/bookingapp/internal/config"
	"github.com/paulbyrne/bookingapp/internal/handlers"
	"github.com/paulbyrne/bookingapp/internal/render"
)

const port = ":8080"

var appConfig config.AppConfig
var session *scs.SessionManager

// main the the main function which is the entry point to the web application
func main() {

	appConfig.InProduction = false

	// session config
	session = scs.New()
	session.Lifetime = 45 * time.Minute
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = appConfig.InProduction

	// set session in appConfig
	appConfig.Session = session

	// create web template
	tempCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	// update template cache in appConfig
	appConfig.TemplateCache = tempCache
	appConfig.UseCache = false

	repo := handlers.NewRepo(&appConfig)
	handlers.NewHandlers(repo)

	render.NewTemplates(&appConfig)

	fmt.Println(fmt.Sprintf("Starting server on %s | http://localhost:8080", port))

	serve := &http.Server{
		Addr:    port,
		Handler: routes(&appConfig),
	}

	err = serve.ListenAndServe()
	log.Fatal(err)
}
