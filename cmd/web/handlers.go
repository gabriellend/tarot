package main

import (
	"html/template"
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("./ui/html/home.page.tmpl")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
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
