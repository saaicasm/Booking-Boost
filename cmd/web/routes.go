package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/iamlego/bookingBoost/pkg/config"
	handler "github.com/iamlego/bookingBoost/pkg/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	// mux := pat.New()

	// mux.Get("/", http.HandlerFunc(handler.Repo.Home))
	// mux.Get("/About", http.HandlerFunc(handler.Repo.About))
	mux := chi.NewRouter()

	mux.Use(WriteToDConsole)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handler.Repo.Home)
	mux.Get("/About", handler.Repo.About)

	fileServer := http.FileServer(http.Dir("./static"))
	log.Println("The content of file server are : ", fileServer) // it is &{./static}

	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
