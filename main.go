package main

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	fmt.Fprintf(w, "<h1>test site</h1>")
}

func contactHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	fmt.Fprintf(w, "<h1>Contact page</h1><p>to get in contact use <a href=\"mailto:mike@dubdubdub.co.uk\">mike@dubdubdub.co.uk</a>")
}

func faqHandler(w http.ResponseWriter, r *http.Request)  {
	w.Header().Set("Content-Type", "text/html, charset=utf-8")
	fmt.Fprintf(w, `
<h1>FAQ page</h1>
<ul>
	<li>Q: question 1</li>
	<li>A: answer 1</li>
</ul>
	`)
}

func main()  {

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", homeHandler)
	r.Get("/contact", contactHandler)
	r.Get("/faq", faqHandler)
	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", r)
}