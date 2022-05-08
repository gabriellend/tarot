package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/cards", displayCards)
	mux.HandleFunc("/cards/major", displayMajor)
	mux.HandleFunc("/cards/minor", displayMinor)
	mux.HandleFunc("/cards/wands", displayWands)
	mux.HandleFunc("/cards/cups", displayCups)
	mux.HandleFunc("/cards/swords", displaySwords)
	mux.HandleFunc("/cards/pentacles", displayPentacles)

	log.Println("Starting server on port 8000")
	err := http.ListenAndServe(":8000", mux)
	log.Fatal(err)
}
