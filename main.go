package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title       string
	Name        string
	Description string
}

// ---------------Page Handlers---------------
func homepageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Handler function called")
	tmpl, err := template.ParseFiles("templates/base.html", "templates/index.html")
	if err != nil {
		log.Println("Error parsing templates:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	data := PageData{
		Title:       "Slurp",
		Name:        "Kelley & Sergio",
		Description: "A brief description about the team behind Slurp.",
	}
	err = tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Println("Error executing template:", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
	}
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("You clicked the button")
	fmt.Fprintln(w, "Hello from the server!")
}

func main() {
	//---------------Routers---------------
	http.HandleFunc("/", homepageHandler)
	http.HandleFunc("/message", messageHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	//---------------Start Server---------------
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
