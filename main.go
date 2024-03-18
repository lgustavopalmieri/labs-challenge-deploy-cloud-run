package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/lgustavopalmieri/labs-challenge-deploy-cloud-run/viacepapi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/temperature/{cep}", viacepapi.HandleTemp)

	log.Println("Servidor iniciado")
	http.ListenAndServe(":8080", r)
}
