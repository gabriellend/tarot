package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	addr := flag.String("addr", ":8000", "HTTP network address")
	flag.Parse()

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

	log.Printf("Starting server on port %s\n", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
