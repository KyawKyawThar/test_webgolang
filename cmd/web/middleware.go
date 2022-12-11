package main

import (
	"fmt"
	"github.com/justinas/nosurf"
	"net/http"
)

func WriteToTestMiddleware(next http.Handler) http.Handler {
	// Handler responds to an HTTP request

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")

		next.ServeHTTP(w, r)
	})
}

// Nosurf add CSRF protection to all post request
func Nosurf(next http.Handler) http.Handler {

	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		Path:     "/",
		Secure:   app.InProduction,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler

}


func SessionLoad(next http.Handler)http.Handler{
	return session.LoadAndSave(next)
}
