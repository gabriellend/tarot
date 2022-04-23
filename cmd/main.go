package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Would you like your fortune told?"))
}

func displayCards(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The whole deck."))
}

func displayMajor(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The major arcana."))
}

func displayMinor(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The minor arcana."))
}

func displayWands(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The wands."))
}

func displayCups(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The cups."))
}

func displaySwords(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The swords."))
}

func displayPentacles(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The pentacles."))
}

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
