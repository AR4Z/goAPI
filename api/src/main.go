package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/rs/cors"

	"./manager"
	"./resources"
)

func main() {
	dg := manager.DgraphClient()
	manager.Setup(dg)

	// router setup
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	// resources routes
	router.Get("/sync", resources.Sync)
	router.Get("/buyers", resources.BuyersCollection)
	router.Get("/buyers/{buyerId}", resources.BuyersDetail)

	httpHandler := cors.Default().Handler(router)

	http.ListenAndServe(":3333", httpHandler)
}
