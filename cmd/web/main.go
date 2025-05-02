package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ZackDiego/bookings/pkg/config"
	"github.com/ZackDiego/bookings/pkg/handlers"
	"github.com/ZackDiego/bookings/pkg/render"
	"github.com/alexedwards/scs/v2"
)


var visitorCount int = 0

const portNumber = ":8080"



func VisitorCount() int{
	return visitorCount
}

var app config.AppConfig
var session *scs.SessionManager

func main() {
	
	// Change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session


	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("error creating template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	app.Session = session

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	// http.HandleFunc("/",  handlers.Repo.Home)
	// http.HandleFunc("/about", handlers.Repo.About)
	// http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
	// 	// Do nothing or serve an icon if you like
	// })
	// _ = http.ListenAndServe(portNumber, nil)
	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	log.Println("Starting server on port " + portNumber)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
