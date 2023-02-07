package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/aymenhta/quitter/config"
	"github.com/aymenhta/quitter/handlers"
	"github.com/aymenhta/quitter/helpers"
	"github.com/go-chi/chi/v5"
)

func main() {
	dbUrl := helpers.GetEnvVariable("DATABASE_URL")

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
	r := chi.NewRouter()

	r.Use(config.LogRequest)

	// This is just a test endpoint
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
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

		helpers.EncodeRes(w, posts)
	})

	// AUTH
	r.Route("/auth", func(r chi.Router) {
		r.Post("/signup", handlers.SignUp)
		r.Post("/signin", handlers.SignIn)
	})

	// POSTS
	r.Route("/posts", func(r chi.Router) {
		r.Get("/", handlers.GetPosts)
		r.Post("/", handlers.CreatePost)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", handlers.PostDetails)
			r.Delete("/", handlers.DeletePost)
		})

		r.Get("/by-user/{userId}", handlers.GetUserPosts)
	})

	return r
}
