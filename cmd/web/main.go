package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

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

	fmt.Printf("Starting Server on %s\n", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
