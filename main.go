package main

import (
	"log"
	"net/http"
)

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
