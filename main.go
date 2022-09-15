package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Sachin47singh/go-course/pkg/config"
	"github.com/Sachin47singh/go-course/pkg/handlers"
	"github.com/Sachin47singh/go-course/pkg/render"
	"github.com/alexedwards/scs/v2"
)

// const portNumber = ":8080"

// // main is the main function
// func main() {
// 	http.HandleFunc("/", handlers.Home)
// 	http.HandleFunc("/about", handlers.About)

//		fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
//		_ = http.ListenAndServe(portNumber, nil)
//	}
// const portNumber = ":1011"

// main is the main function
// func main() {
// 	var app config.AppConfig

// 	tc, err := render.CreateTemplateCache()
// 	if err != nil {
// 		log.Fatal("cannot create template cache")
// 	}

// 	app.TemplateCache = tc
// 	app.UseCache = false

// 	repo := handlers.NewRepo(&app)
// 	handlers.NewHandlers(repo)

// 	render.NewTemplates(&app)

// 	// http.HandleFunc("/", handlers.Repo.Home)
// 	// http.HandleFunc("/about", handlers.Repo.About)

//		fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))
//		// _ = http.ListenAndServe(portNumber, nil)
//		srv := &http.Server{
//			Addr:    portNumber,
//			Handler: routes(&app),
//		}
//		err = srv.ListenAndServe()
//		if err != nil {
//			log.Fatal(err)
//		}
//	}
const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

// main is the main function
func main() {
	// change this to true when in production
	app.InProduction = false

	// set up the session
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
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Println(fmt.Sprintf("Staring application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
