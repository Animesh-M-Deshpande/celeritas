package celeritas

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (c *Celeritas) routes() http.Handler {
	//cheat router is being used in favour to standard router for go

	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)

	if c.Debug {
		mux.Use(middleware.Logger)
	}

	mux.Use(middleware.Recoverer) //recover if the application panics

	mux.Get("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to Celeritas")
	})
	return mux
}
