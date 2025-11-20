package main

import (
	"log/slog"
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

	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	slog.SetDefault(logger)

	err := api.run(api.mount())
	if err != nil {
		slog.Error("could not start server", "error", err)
		os.Exit(1)
	}
}
