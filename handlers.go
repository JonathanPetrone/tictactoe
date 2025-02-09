package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func (t *Tictactoe) renderTemplate(w http.ResponseWriter, dynamicOnly bool) {
	var err error
	if dynamicOnly {
		// only updates dynamic_content
		err = DynamicContentTmpl.Execute(w, t)
	} else {
		// updates all
		err = Tmpl.Execute(w, t)
	}

	if err != nil {
		log.Printf("Template execution error: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func (t *Tictactoe) ServeStart(w http.ResponseWriter, r *http.Request) {
	t.renderTemplate(w, false)
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
	x, errX := strconv.Atoi(r.URL.Query().Get("x"))
	y, errY := strconv.Atoi(r.URL.Query().Get("y"))

	if errX != nil || errY != nil {
		http.Error(w, "Invalid move coordinates", http.StatusBadRequest)
		return
	}

	// Ensure the move is valid
	if x < 0 || x >= 3 || y < 0 || y >= 3 || t.Board[x][y] != "" {
		http.Error(w, "Invalid move", http.StatusBadRequest)
		return
	}

	// Make the move
	t.Board[x][y] = t.PlayerTurn

	// Check if the player has won
	if t.checkWin(t.PlayerTurn) {
		t.GameStatus = fmt.Sprintf("Player %s wins!", t.PlayerTurn)
	} else {
		// Switch turns
		t.switchTurn()
		t.GameStatus = fmt.Sprintf("It's %s's turn", t.PlayerTurn)
	}

	if t.TurnNumber == 9 {
		if t.checkWin(t.PlayerTurn) {
			t.GameStatus = fmt.Sprintf("Player %s wins!", t.PlayerTurn)
		} else {
			t.GameStatus = fmt.Sprintf("It's a draw")
		}
	}
	t.TurnNumber++

	t.renderTemplate(w, true)
}

func (t *Tictactoe) resetBoard(w http.ResponseWriter, r *http.Request) {
	t.Init() // Reset the board to initial state
	t.renderTemplate(w, true)
}
