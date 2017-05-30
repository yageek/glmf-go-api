package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world \n")
}

func main() {
	// Source : https://golang.org/pkg/net/http/
	mux := http.NewServeMux()
	mux.Handle("/hello/", http.HandlerFunc(helloHandler))
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		// Le motif "/" correspond à n’importe quel chemin, c’est
		// pourquoi il faut vérifier si la requête correspond au chemin racine.
		if req.URL.Path != "/" {
			http.NotFound(w, req)
			return
		}
		fmt.Fprintf(w, "Bienvenue sur la page racine!")
	})
	http.ListenAndServe(":8080", mux)
}
