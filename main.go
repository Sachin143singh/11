package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Sachin47singh/go-course/pkg/config"
	"github.com/Sachin47singh/go-course/pkg/handlers"
	"github.com/Sachin47singh/go-course/pkg/render"
)

// const portNumber = ":8080"

// // main is the main function
// func main() {
// 	http.HandleFunc("/", handlers.Home)
// 	http.HandleFunc("/about", handlers.About)

// 	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
// 	_ = http.ListenAndServe(portNumber, nil)
// }
const portNumber = ":1010"

// main is the main function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}

