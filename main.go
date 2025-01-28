package main

import (
	"fmt"
	"net/http"
)

// Define a home handler function which writes a byte
// slice containing "Hello from snnipetbox" as the response body

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	
	w.Write([]byte("Hello from snippetbox"))
}

func showSnnipet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Display a specific snippet..."))
}

func createSnnipet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}
	w.Write([]byte("Create a new snippet..."))
}
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/", showSnnipet)
	mux.HandleFunc("/snippet/create/", createSnnipet)


	fmt.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)

	fmt.Println(err)
}
