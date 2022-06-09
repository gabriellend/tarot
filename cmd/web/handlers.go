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

	files := []string{
		"./ui/html/home.page.tmpl",
		"./ui/html/base.layout.tmpl",
	}

	ts, err := template.ParseFiles(files...)
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

func showCards(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The whole deck."))
}

func showMajor(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The major arcana."))
}

func showMinor(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The minor arcana."))
}

func showWands(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The wands."))
}

func showCups(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The cups."))
}

func showSwords(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The swords."))
}

func showPentacles(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("The pentacles."))
}
