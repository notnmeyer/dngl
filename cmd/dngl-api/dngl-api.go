package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"

	"github.com/notnmeyer/dngl/internal/api/handler"
	"github.com/notnmeyer/dngl/internal/api/middleware"
	"github.com/notnmeyer/dngl/internal/envhelper"
)

func main() {
	r := chi.NewRouter()

	r.Use(
		chiMiddleware.Logger,
		middleware.ContextInjector,
		middleware.BearerTokenValidation,
	)

	r.Get("/healthcheck", handler.Healthcheck)

	r.Post("/note/create", handler.CreateNote)
	r.Get("/note/{id}", handler.GetNote)
	r.Post("/note/delete/{id}", handler.DeleteNote)
	r.Get("/notes", handler.ListNotes)

	addr := fmt.Sprintf(":%s", envhelper.New().DNGL_API_PORT)
	fmt.Printf("starting server on %s...\n", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
