package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/stackriv/go-api-starter/internal/middleware"
	"github.com/stackriv/go-api-starter/internal/pkg"
)

func main() {
	// signals channel
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM)

	// Goroutine to handle signals
	go func() {
		<-signals
		fmt.Println("\nReceived an interrupt, stopping services...")
		os.Exit(0)
	}()

	err := start(os.Args[1:])
	if err != nil {
		log.Println("Error starting the application:", err)
		os.Exit(0)
	}
	select {}
}

func start(args []string) error {
	if len(args) != 0 {
		return errors.New("no arguments are supported")
	}

	// Check if .env file exists
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		return errors.New(".env file not found")
	}

	// Read env file
	err := pkg.Env()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux = routes(mux)

	wrappedMux := middleware.LoggingMiddleware(mux)
	wrappedMux = middleware.CORSMiddleware(wrappedMux)
	wrappedMux = middleware.ErrorMiddleware(wrappedMux)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", os.Getenv("APP_PORT")),
		Handler: wrappedMux,
	}

	log.Println("Server is listening on", os.Getenv("APP_PORT"))
	return server.ListenAndServe()
}

func routes(mux *http.ServeMux) *http.ServeMux {
	// mux = userController.UserRoutes(mux)
	return mux
}
