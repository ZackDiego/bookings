package main

import (
	"net/http"
	"log"
	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

// NoSurf adds CSRF protection to all POST requests
func NoSurf(next http.Handler) http.Handler {
	csrfHandler := nosurf.New(next)
	csrfHandler.SetBaseCookie(http.Cookie{
		Path: "/",
		HttpOnly: true,
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})
	return csrfHandler
}

// SessionLoad loads and saves the session for every request
func SessionLoad(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}
	