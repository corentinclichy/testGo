package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/corentinclichy/myapp/pkg/config"
	"github.com/corentinclichy/myapp/pkg/handlers"
	"github.com/corentinclichy/myapp/pkg/render"
)

const portNumber = ":8080"

// @Title main
// @Description this is the main application function
func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache", err)
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)

	handlers.NewHandler(repo)
	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)
}
