package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"]
		page := vars["page"]

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)

	})

	r.HandleFunc("/books/{title}", getBook).Methods("GET")
	r.HandleFunc("/books/{title}", postBook).Methods("POST")

	port := "5100"

	http.ListenAndServe(":"+port, r)
	fmt.Printf("Running server at port %s", port)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "GET request for book with title %s\n", vars["title"])
}

func postBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprintf(w, "POST request for book with title %s\n", vars["title"])
}
