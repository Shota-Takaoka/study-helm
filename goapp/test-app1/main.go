package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/caarlos0/env"
)

type ServerConfig struct {
	Port string `env:"PORT" envDefault:"8080"`
}

func main() {
	var cfg ServerConfig
	if err := env.Parse(&cfg); err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("requested host=%s", r.Host)
		fmt.Fprintf(w, "test-app1\n")
	})

	log.Printf("Starting server on port %s", cfg.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", cfg.Port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
