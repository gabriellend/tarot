package main

import (
	"fmt"

	"github.com/gabriellend/tarot/cmd/tarot/config"
)

func main() {
	// cfg := config.Load()

	// srv, err := server.New(server.Params{
	// 	Debug: cfg.Debug,
	// })
	// if err != nil {
	// 	log.Fatalf("failed to create server: %v", err)
	// }

	// fmt.Printf("%T\n%#v\n", srv, srv)

	fmt.Println("starting")

	// Setup the main context.
	// ctx, cancel := context.WithCancel(context.Background())

	cfg, err := config.New()
	if err != nil {
		fmt.Printf("failed to get config: %v\n", err)
	}

	fmt.Printf("%+v", cfg)

	// server.New()
}
