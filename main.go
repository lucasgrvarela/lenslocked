package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/go-chi/chi/v5"
	"github.com/lucasgrvarela/lenslocked/controllers"
	"github.com/lucasgrvarela/lenslocked/views"
)

func main() {
	r := chi.NewRouter()

	tmpl, err := views.Parse(filepath.Join("templates", "home.tmpl"))
	if err != nil {
		panic(err)
	}
	r.Get("/", controllers.StaticHandler(tmpl))

	tmpl, err = views.Parse(filepath.Join("templates", "contact.tmpl"))
	if err != nil {
		panic(err)
	}
	r.Get("/contact", controllers.StaticHandler(tmpl))

	tmpl, err = views.Parse(filepath.Join("templates", "faq.tmpl"))
	if err != nil {
		panic(err)
	}
	r.Get("/faq", controllers.StaticHandler(tmpl))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
	})

	fmt.Println("Starting the server on :3000...")
	http.ListenAndServe(":3000", r)
}
