package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	port := "3333"

	// router setup
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// resources routes
	router.Get("/sync", resources.sync)
	router.Get("/buyers", resources.buyersCollection)
	router.Get("/buyers/{buyerId}", resources.buyersDetail)

}
