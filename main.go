package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Shubham-Khetan-2005/rssagg/internal/database"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

// main runs the HTTP server
func main() {
	godotenv.Load()

	// The server is configured to listen on the PORT environment variable, and
	// connect to the database at the URL specified in DB_URL.

	portString := os.Getenv("PORT")
	if portString == "" {
		log.Fatal("PORT must be set")
	}

	dbURL := os.Getenv("DB_URL")
	if portString == "" {
		log.Fatal("DB_URL not found in env")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
	}

	queries := database.New(conn)

	apiCfg := apiConfig{
		DB: queries,
	}

	go startScrapping(queries, 10, time.Minute)

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	// The server responds to the following endpoints:
	//
	// POST /v1/users: Creates a new user with a random API key.
	// GET /v1/users: Returns the currently authenticated user.
	//
	// POST /v1/feeds: Creates a new feed with the given name and URL.
	// GET /v1/feeds: Returns all feeds of all users.
	//
	// POST /v1/feed_follows: Creates a new feed follow for the currently
	// authenticated user.
	// GET /v1/feed_follows: Returns all feed follows for the currently
	// authenticated user.
	// DELETE /v1/feed_follows/{feedFollowID}: Deletes the feed follow with the
	// given ID for the currently authenticated user.


	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/error", handlerErr)

	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

	v1Router.Post("/feeds", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
	v1Router.Get("/feeds", apiCfg.handlerGetFeeds)

	v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetPostsForUser))

	v1Router.Post("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollow))
	v1Router.Get("/feed_follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
	v1Router.Delete("/feed_follows/{feedFollowID}", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollow))

	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Addr:    ":" + portString,
		Handler: router,
	}

	fmt.Printf("Listening on port %s\n", portString)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
