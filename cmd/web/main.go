package main

import (
	"fmt"
	"time"

	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/anubhavjoshi040/go-webapp-template/config"
	"github.com/anubhavjoshi040/go-webapp-template/pkg/handlers"
	"github.com/anubhavjoshi040/go-webapp-template/pkg/render"
)

const portNumber = ":4040"
var app config.AppConfig
var session *scs.SessionManager

// main is the main application function
func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("starting application on port %s", portNumber))

	srv := &http.Server {
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}