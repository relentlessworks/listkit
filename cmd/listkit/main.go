package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/relentlessworks/listkit/internal/api"
	"github.com/relentlessworks/listkit/internal/config"
)

func main() {
	cfg := config.Load()

	handler := api.NewHandler(cfg.NoAuth)
	mux := handler.Routes()

	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: mux,
	}

	// Graceful shutdown
	go func() {
		sigCh := make(chan os.Signal, 1)
		signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
		<-sigCh
		log.Println("shutting down...")
		server.Close()
	}()

	fmt.Fprintf(os.Stderr, "listkit listening on %s\n", cfg.Addr)
	if cfg.NoAuth {
		fmt.Fprintf(os.Stderr, "auth disabled (development mode)\n")
	}
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server error: %v", err)
	}
}
