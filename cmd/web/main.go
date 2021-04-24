package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"

	"github.com/danarcheronline/wizard-bookings/pkg/config"
	"github.com/danarcheronline/wizard-bookings/pkg/render"

	"github.com/danarcheronline/wizard-bookings/pkg/handlers"
)

var app config.AppConfig
var session *scs.SessionManager

const portNumber = ":8080"

func main() {

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app.TemplateCach = tc
	app.UseCache = false

	render.NewTemplates(&app)

	fmt.Printf("Starting application in port %s", portNumber)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
