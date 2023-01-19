package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("Hello, gophers;"))
	})

	log.Println("App started on port :4000")
	err := http.ListenAndServe(":4000", logRequest(mux))
	if err != nil {
		log.Fatal(err)
	}
}
