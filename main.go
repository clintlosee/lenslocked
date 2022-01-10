package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
// homeView    *views.View
// contactView *views.View
// signupView  *views.View
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Welcome to my awesome site!</h1>")
	// must(homeView.Render(w, nil), w)
	// tplPath := filepath.Join("templates", "home.gohtml")
	// executeTemplate(w, tplPath)
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>Contact Page</h1><p>To get in touch, email me here -></p>")
	// must(contactView.Render(w, nil), w)
}

func faqHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "<h1>FAQ Page</h1><p>Here are some frequently asked questions</p>")
	// must(contactView.Render(w, nil), w)
}

// func signupHandler(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	must(signupView.Render(w, nil), w)
// }

func main() {
	// homeView = views.NewView("bootstrap", "views/home.gohtml")
	// contactView = views.NewView("bootstrap", "views/contact.gohtml")
	// signupView = views.NewView("bootstrap", "views/signup.gohtml")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	// r.Get("/signup", signupHandler)
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Page not found", http.StatusNotFound)
	})
	fmt.Println("Starting the server on :7000...")
	http.ListenAndServe(":7000", r)

}

// func must(err error, w http.ResponseWriter) {
// 	if err != nil {
// 		// panic(err)
// 		log.Printf("parsing template: %v", err)
// 		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
// 		return
// 	}
// }

// func executeTemplate(w http.ResponseWriter, filepath string) {
// 	t, err := views.Parse(filepath)
// 	if err != nil {
// 		log.Printf("parsing template: %v", err)
// 		http.Error(w, "There was an error parsing the template.", http.StatusInternalServerError)
// 		return
// 	}
// 	t.Execute(w, nil)
// }
