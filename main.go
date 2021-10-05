package main

import (
	"fmt"
	"net/http"

	"github.com/clintlosee/lenslocked/views"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	homeView    *views.View
	contactView *views.View
	signupView  *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(contactView.Render(w, nil))
}

func signup(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(signupView.Render(w, nil))
}

func main() {
	homeView = views.NewView("bootstrap", "views/home.gohtml")
	contactView = views.NewView("bootstrap", "views/contact.gohtml")
	signupView = views.NewView("bootstrap", "views/signup.gohtml")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", home)
	r.Get("/contact", contact)
	r.Get("/signup", signup)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :7000...")
	http.ListenAndServe(":7000", r)

	// r := mux.NewRouter()
	// r.HandleFunc("/", home)
	// r.HandleFunc("/contact", contact)
	// r.HandleFunc("/signup", signup)
	// http.ListenAndServe(":7000", r)
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
