package main

import (
	"ejemplo-multi-contenedor-docker-compose/controllers"
	"ejemplo-multi-contenedor-docker-compose/models"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	handler := controllers.New()

	server := &http.Server{
		Addr:    "0.0.0.0:8008",
		Handler: handler,
	}

	models.ConnectDB()

	server.ListenAndServe()
}
