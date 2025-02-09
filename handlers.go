package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func (t *Tictactoe) ServeStart(rw http.ResponseWriter, r *http.Request) {
	// Parse the template file
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(rw, "Error loading template", http.StatusInternalServerError)
		log.Println("Template parsing error:", err)
		return
	}

	// Execute the template with the current game state
	err = tmpl.Execute(rw, t)
	if err != nil {
		http.Error(rw, "Error rendering template", http.StatusInternalServerError)
		log.Println("Template execution error:", err)
	}
}

func (t *Tictactoe) switchTurn() {
	if t.PlayerTurn == "X" {
		t.PlayerTurn = "O"
	} else {
		t.PlayerTurn = "X"
	}
}

func (t *Tictactoe) makeMoveHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the coordinates from the request
	x, _ := strconv.Atoi(r.URL.Query().Get("x"))
	y, _ := strconv.Atoi(r.URL.Query().Get("y"))

	// Make sure the square is empty
	if t.Board[x][y] == "" {
		// Set the move on the board
		t.Board[x][y] = t.PlayerTurn
		// Switch turns
		t.switchTurn()
	}

	// Respond with the updated move (X or O)
	fmt.Fprintf(w, "%s", t.Board[x][y])
	t.ServeStart(w, r)
}

func (t *Tictactoe) resetBoard(w http.ResponseWriter, r *http.Request) {
	t.Init()
	t.ServeStart(w, r)
}
