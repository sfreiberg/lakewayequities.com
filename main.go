package main

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//go:embed public/*
var staticFiles embed.FS

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Create a sub filesystem from the embedded files
	publicFS, err := fs.Sub(staticFiles, "public")
	if err != nil {
		log.Fatal(err)
	}

	fileServer := http.FileServer(http.FS(publicFS))
	r.Handle("/*", fileServer)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}