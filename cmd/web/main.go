package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	// create a file server which serves files out of the
	// "./ui/static"directory
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server
	// as the handler for all URL paths that start with "/static"
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	fmt.Println("http://localhost:4000")
	fmt.Println("Starting Server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
