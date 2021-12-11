package main

import (
	"fmt"
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

func main()  {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/contact", contactHandler)

	fmt.Println("Starting server on :3000...")
	http.ListenAndServe(":3000", nil)
}