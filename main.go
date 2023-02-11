package main

import (
	"log"
	"net/http"
)

// Homepage handler
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hi from Snippetbox"))
}

// handler to show the content of the note
func showSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Showing notes..."))
}

// handler to create a new note
func createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Method Not Allowed!", 405)
		return
	}

	w.Write([]byte("Creation of a new note..."))
}

func main() {
	/*
		---------------------------------------------------
		IMPORTANT!
		Instead of using http.HandleFunc use NewServeMux to create a var locally

		http registers its paths using DefaultServeMux which is saved as a global variable in net/http
		therefore making app vulnerable to external packages being written and getting access to the app
		---------------------------------------------------
	*/
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Listening on http://127.0.0.1:4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

/*
continuation: https://golangify.com/url-query-strings
*/
