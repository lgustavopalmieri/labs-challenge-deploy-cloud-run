package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	_ "github.com/lgustavopalmieri/labs-challenge-deploy-cloud-run/docs"
	"github.com/lgustavopalmieri/labs-challenge-deploy-cloud-run/viacepapi"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title           Labs challenge: Deploy Cloud Run
// @version         1.0
// @description     This API is designed to provide a seamless integration of location and weather information 

// @license.name   Luiz Palmieri

// @host            localhost:8080
// @BasePath        /
func main() {
	r := chi.NewRouter()
	r.Get("/temperature/{cep}", viacepapi.HandleTemp)

	r.Get("/docs/*", httpSwagger.Handler(httpSwagger.URL("http://localhost:8080/docs/doc.json")))
	log.Println("Servidor iniciado")
	http.ListenAndServe(":8080", r)
}
