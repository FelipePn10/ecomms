package main

import (
	"log"
	"os"
)

func main() {
	cfg := config{
		addr: ":6800",
		db:   dbConfig{},
	}

	api := application{
		config: cfg,
	}

	err := api.run(api.mount())
	if err != nil {
		log.Println("Server error:", err)
		os.Exit(1)
	}
}
