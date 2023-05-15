package controllers

import (
  "net/http"
  "github.com/gorilla/mux"
)


func New() http.Handler {
  router := mux.NewRouter()

  // Definimos las rutas

  router.HandleFunc("/bluetabers",getBluetabers).Methods("GET")
  router.HandleFunc("/bluetabers/{id}",getBluetaber).Methods("GET")
  
  return router
}
