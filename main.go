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

// Parse all templates
// var tmpl = template.Must(template.ParseGlob("templates/*.html"))

// var tmpl *template.Template

// // Initialize templates by parsing all HTML files once
// func initTemplates() {
//     var err error
//     tmpl, err = template.ParseGlob("templates/*.html")
//     if err != nil {
//         log.Fatalf("Error parsing templates: %v", err)
//     }
// }

// Render template with specified content template (e.g., "about.html")
func renderTemplate(w http.ResponseWriter, tmplName string, data PageData) {
	baseTemplate := "templates/base.html"
    pageTemplate := fmt.Sprintf("templates/%s", tmplName)
	tmpl, err := template.ParseFiles(baseTemplate, pageTemplate)
    
	err = tmpl.ExecuteTemplate(w, "base", data)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
}

func homepageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Home handler function called")
	data := PageData{
		Title:       "Slurp",
		Name:        "Kelley & Sergio",
		Description: "A brief description about the team behind Slurp.",
	}
	// tmpl, err := template.ParseFiles("base.html", "index.html")
	// err := tmpl.ExecuteTemplate(w, "base", data)
	// if err != nil {
	// 	log.Println("Error executing template:", err)
	// 	http.Error(w, "Internal server error", http.StatusInternalServerError)
	// }
	renderTemplate(w, "index.html", data)
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("You clicked the button")
	fmt.Fprintln(w, "Hello from the server!")
}

func kelleyHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Kelley handler function called")

	data := PageData{
		Title: "About Kelley",
	}
	// err := tmpl.ExecuteTemplate(w, "base", data)
	// if err != nil {
	// 	log.Println("Error executing template for Kelley:", err)
	// 	http.Error(w, "Internal server error", http.StatusInternalServerError)
	// }
	renderTemplate(w, "kelley.html", data)
}

func sergioHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Sergio handler function called")
	data := PageData{
		Title: "About Sergio",
	}
	renderTemplate(w, "sergio.html", data)
	// err := tmpl.ExecuteTemplate(w, "templates/sergio.html", data)
	// if err != nil {
	// 	log.Println("Error executing template for Sergio:", err)
	// 	http.Error(w, "Internal server error", http.StatusInternalServerError)
	// }
}

func main() {
	//---------------Routers---------------
	http.HandleFunc("/message", messageHandler)
	http.HandleFunc("/kelley", kelleyHandler)
    http.HandleFunc("/sergio", sergioHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", homepageHandler)
	//---------------Start Server---------------
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
