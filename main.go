package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/clintlosee/lenslocked/controllers"
	"github.com/clintlosee/lenslocked/views"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "home.gohtml")))))

	r.Get("/contact", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "contact.gohtml")))))

	r.Get("/faq", controllers.StaticHandler(
		views.Must(views.Parse(filepath.Join("templates", "faq.gohtml")))))

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :7000...")
	http.ListenAndServe(":7000", r)
}
