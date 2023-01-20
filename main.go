package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aymenhta/quitter/config"
	"github.com/aymenhta/quitter/handlers"
)

func main() {
	dbUrl := getEnvVariable("DATABASE_URL")

	// Config default logging
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Connect to db
	db, err := config.SetupDb(dbUrl)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()
	config.TestDb(db)

	// Init app
	config.InitApplication(db, infoLog, errorLog)

	// Run the server
	infoLog.Println("App started on port :4000")
	err = http.ListenAndServe(":4000", routes())
	if err != nil {
		errorLog.Fatal(err)
	}
}

func routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		w.Write([]byte("Hello, gophers;"))
	})

	mux.HandleFunc("/users/signup", handlers.SignUp)

	return logRequest(mux)
}
