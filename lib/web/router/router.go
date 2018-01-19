package router

import (
	"net/http"

	"github.com/JamesClonk/iRnotify/lib/web/html"
	"github.com/gorilla/mux"
)

func New() *mux.Router {
	router := mux.NewRouter()
	setupRoutes(router)
	return router
}

func setupRoutes(router *mux.Router) *mux.Router {
	// HTML
	router.NotFoundHandler = http.HandlerFunc(html.NotFound)

	router.HandleFunc("/", html.Index)
	router.HandleFunc("/error", html.ErrorHandler)

	router.HandleFunc("/racers", html.Racers)
	router.HandleFunc("/racer/{id}", html.Racer)

	return router
}
