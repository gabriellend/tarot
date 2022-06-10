package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/cards", showCards)
	mux.HandleFunc("/cards/major", showMajor)
	mux.HandleFunc("/cards/minor", showMinor)
	mux.HandleFunc("/cards/wands", showWands)
	mux.HandleFunc("/cards/cups", showCups)
	mux.HandleFunc("/cards/swords", showSwords)
	mux.HandleFunc("/cards/pentacles", showPentacles)

	// Serve files
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on port 8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
