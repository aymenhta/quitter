package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aymenhta/quitter/config"
	"github.com/aymenhta/quitter/handlers"
	"github.com/julienschmidt/httprouter"
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
	server := http.Server{
		Addr:         ":4000",
		ErrorLog:     errorLog,
		Handler:      routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	infoLog.Println("SERVER started on port :4000")
	err = server.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}

func routes() http.Handler {
	router := httprouter.New()

	// This is just a test endpoint
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		type Post struct {
			Id       int       `json:"id"`
			Content  string    `json:"content"`
			PostedAt time.Time `json:"postedAt"`
		}

		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		posts := []*Post{
			{Id: 1, Content: "post #1", PostedAt: time.Now()},
			{Id: 2, Content: "post #2", PostedAt: time.Now()},
			{Id: 3, Content: "post #3", PostedAt: time.Now()},
			{Id: 4, Content: "post #4", PostedAt: time.Now()},
		}

		// marshal
		encoder := json.NewEncoder(w)
		err := encoder.Encode(posts)
		if err != nil {
			http.Error(w, "Could not marshal json", http.StatusInternalServerError)
			return
		}
	})

	// AUTH
	router.HandlerFunc(http.MethodPost, "/auth/signup", handlers.SignUp)
	router.HandlerFunc(http.MethodPost, "/auth/signin", handlers.SignIn)

	// POSTS
	router.HandlerFunc(http.MethodPost, "/posts", handlers.CreatePost)
	router.HandlerFunc(http.MethodGet, "/posts/:id", handlers.PostDetails)

	return config.LogRequest(router)
}
