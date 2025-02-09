package main

import (
	"html/template"
	"log"
	"net/http"
)

var Tmpl *template.Template
var DynamicContentTmpl *template.Template

func init() {
	var err error
	Tmpl, err = template.ParseFiles("templates/index.html", "templates/dynamic_content.html")
	if err != nil {
		log.Fatalf("Failed to load templates: %v", err)
	}

	DynamicContentTmpl, err = template.ParseFiles("templates/dynamic_content.html")
	if err != nil {
		log.Fatalf("Failed to load dynamic content template: %v", err)
	}
}

func main() {
	board := Tictactoe{}
	board.Init()

	mux := http.NewServeMux()
	server := &http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	mux.Handle("GET /", http.HandlerFunc(board.ServeStart))
	mux.Handle("POST /move", http.HandlerFunc(board.makeMoveHandler))
	mux.Handle("POST /reset", http.HandlerFunc(board.resetBoard))

	err := server.ListenAndServe()

	if err != nil {
		log.Printf("Server error: %v", err)
		log.Fatal(err)
	}
}
