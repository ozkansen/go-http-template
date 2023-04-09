package main

import (
	"log"
	"net/http"
	"time"

	"go-http-template/pkg/api"
	"go-http-template/pkg/logger"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

func main() {
	l := logger.NewDefaultConsoleLogger()
	l.Info().Msg("Starting HTTP API Server")
	err := Run(l)
	if err != nil {
		l.Fatal().Err(err).Send()
	}
}

func Run(l *zerolog.Logger) error {
	app := NewExampleAPI()
	app.SetLogger(l)

	router := api.NewRouter()

	// global middleware tanımlamaları
	router.Use(api.MiddlewareRequestID) // %10 performance decrease
	router.Use(api.MiddlewareLogger(l)) // %75 performance decrease
	router.Use(middleware.Recoverer)

	// uygulamaları ekle
	router.Add("/", app)

	mux, err := router.Router()
	if err != nil {
		return err
	}

	server := &http.Server{
		Addr:              ":8080",
		Handler:           mux,
		ReadTimeout:       1 * time.Second,
		WriteTimeout:      1 * time.Second,
		IdleTimeout:       30 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
		ErrorLog:          log.New(l, "", 0), // TODO: route err channel
	}
	err = server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
