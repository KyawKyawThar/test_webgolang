package main

import (
	"fmt"
	"github.com/KyawKyawThar/gowebtest/pkg/config"
	"github.com/KyawKyawThar/gowebtest/pkg/handlers" //according to go module GitHub
	"github.com/KyawKyawThar/gowebtest/pkg/render"
	"github.com/alexedwards/scs/v2"
	"log"
	"net/http"
	"time"
)

//
//func Home(w http.ResponseWriter, r *http.Request) {
//	b, err := fmt.Fprintf(w, "Helo Golang")
//	if err != nil {
//		return
//	}
//
//	fmt.Println(fmt.Sprintf("Number of byte written %d", b))
//}
//
//func About(w http.ResponseWriter, r *http.Request) {
//
//	value := addValue(2, 5)
//	_, _ = fmt.Fprintf(w, "This is about page and 2 + 5 is %d", value)
//
//}
//
//func Divider(w http.ResponseWriter, r *http.Request) {
//
//	x := 20.0
//	y := 0.0
//	result, err := divideValues(x, y)
//
//	if err != nil {
//		_, _ = fmt.Fprintf(w, "can't divided by %f", y)
//		return
//	}
//
//	_, _ = fmt.Fprintf(w, "This is divider page and x / y is: result %f,%f,%f", x, y, result)
//}
//
//func addValue(x, y int) int {
//
//	return x + y
//}
//
//func divideValues(a, b float64) (float64, error) {
//
//	if b <= 0 {
//		err := errors.New("can't not divided by zero")
//		return 0, err
//	}
//
//	divider := a / b
//	return divider, nil
//}
//
//func main() {
//
//	http.HandleFunc("/", Home)
//	http.HandleFunc("/about", About)
//	http.HandleFunc("/divider", Divider)
//
//	fmt.Println(fmt.Sprintf("Server is running on port %s", portNumber))
//
//	_ = http.ListenAndServe(portNumber, nil)
//
//}

const portNumber = ":8080"

var app config.AppConfig //Only For Assign Data Type, so we don't need pointer

var session *scs.SessionManager

func main() {

	app.InProduction = false

	//For session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.UseCache = false
	app.Session = session

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

	app.TemplateCache = tc
	render.NewTemplates(&app)

	repo := handlers.NewRepo(&app)

	handlers.RepoHandler(repo)
	//fmt.Println("APP", app)

	//http.HandleFunc("/", handlers.Repo.Home)
	//http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Server is running on port %s", portNumber))

	//_ = http.ListenAndServe(portNumber, nil)

	srv := http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
