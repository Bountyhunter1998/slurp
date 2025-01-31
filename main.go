package main

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

// Embed templates and static files
//
//go:embed templates/* static/*
var content embed.FS

type PageData struct {
	Title       string
	Name        string
	Description string
}

// Homepage Handler
func homepageHandler(tmpl *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("[%s] %s %s\n", r.RemoteAddr, r.Method, r.URL.Path)

		data := PageData{
			Title:       "Slurp",
			Name:        "Kelley & Sergio",
			Description: "A brief description about the team behind Slurp.",
		}
		err := tmpl.ExecuteTemplate(w, "base", data)
		if err != nil {
			log.Println("Error executing template:", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	}
}

// Message Handler
func messageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("You clicked the button")
	fmt.Fprintln(w, "Hello from the server!")
}

func main() {
	// Parse templates once at startup
	tmpl := template.Must(template.ParseFS(content, "templates/base.html", "templates/index.html"))

	// Serve static files using embedded filesystem
	staticFS := http.FS(content)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(staticFS)))

	// Handle favicon requests to avoid log spam
	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	// Register routes
	http.HandleFunc("/", homepageHandler(tmpl))
	http.HandleFunc("/message", messageHandler)

	// Get port from environment variable (default 8080)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe("0.0.0.0:8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
