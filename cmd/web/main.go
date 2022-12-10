package main

import (
	"fmt"
	"github.com/KyawKyawThar/gowebtest/pkg/config"
	"github.com/KyawKyawThar/gowebtest/pkg/handlers" //according to go module GitHub
	"github.com/KyawKyawThar/gowebtest/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

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

func main() {

	var app config.AppConfig //Only For Assign Data Type and so we don't need pointer

	app.UseCache = false

	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)

	handlers.RepoHandler(repo)
	//fmt.Println("APP", app)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Server is running on port %s", portNumber))

	_ = http.ListenAndServe(portNumber, nil)

}
