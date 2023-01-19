package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, "URL Path Could not be found", 404)
			return
		}
		w.Write([]byte("Hello, gophers;"))
	})

	log.Fatal(http.ListenAndServe(":4000", mux))
}
