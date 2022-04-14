package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gabriellend/tarot/cmd/tarot/config"
	"github.com/gabriellend/tarot/internal/api/templates"
)

// type Server struct {
// 	debug bool //Why is this not capitalized? Cause it's not used anywhere right now?
// }

// type Params struct {
// 	Debug bool
// }

// func printHello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("\nhello")
// }

// func New(p Params) (*Server, error) {
// 	http.HandleFunc("/printHello", printHello)

// 	log.Fatal(http.ListenAndServe(":8080", nil))
// 	fmt.Printf("using port 8080")

// 	return &Server{
// 		debug: p.Debug,
// 	}, nil
// }

func serve(ctx context.Context, done chan struct{}, cfg *config.Config) {
	fmt.Println("using port %s", cfg.Port)

	// Load the templates.
	tmpl, err := templates.Parse(cfg.TemplatesDir)
	if err != nil {
		fmt.Errorf("failed to load templates: %v", err)
	}

	router, err := router.New(router.Params{
		Templates: tmpl,
		StaticDir: cfg.StaticDir,
	})

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: router,
	}

}
