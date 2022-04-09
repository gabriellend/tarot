package server

import (
	"fmt"
	"log"
	"net/http"
)

type Server struct {
	debug bool
}

type Params struct { // Do you always need a Params struct as a middleman? It seems redundant right now.
	Debug bool
}

func printHello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nhello")
}

func New(p Params) (*Server, error) {
	fmt.Println("starting")

	fmt.Printf("using port 8080")

	http.HandleFunc("/printHello", printHello)

	log.Fatal(http.ListenAndServe(":8080", nil))
	return
	// return &Server{
	// 	debug: p.Debug,
	// }, nil
}
