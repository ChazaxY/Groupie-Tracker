package main

import (
	"html/template"
	"log"
	"net/http"
)

type PageData struct {
	Title       string
	WordToGuess string
	Output      string
	IfCorrect   string
}

func indexPageHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", PageData{
		Title: "Main Page",
	})
}

func renderTemplate(w http.ResponseWriter, tmpl string, data PageData) {
	tmplPath := "templates/" + tmpl
	t, err := template.ParseFiles(tmplPath)
	if err != nil {
		log.Println("Error", err)
		http.Error(w, "Error", http.StatusInternalServerError)
		return
	}
	if err := t.Execute(w, data); err != nil {
		http.Error(w, "Error", http.StatusInternalServerError)
	}
}
func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", indexPageHandler)
	log.Println("Server launched on :5510")
	log.Fatal(http.ListenAndServe(":5510", nil))
    
}