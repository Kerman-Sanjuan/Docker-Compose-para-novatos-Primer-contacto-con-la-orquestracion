package controllers

import (
	"ejemplo-multi-contenedor-docker-compose/models"
	"ejemplo-multi-contenedor-docker-compose/utils"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func getBluetabers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

  var bluetabers []models.Bluetaber
  models.DB.Find(&bluetabers)

  json.NewEncoder(w).Encode(bluetabers)
}

func getBluetaber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

  id := mux.Vars(r)["id"]
  var bluetaber models.Bluetaber

  val := models.DB.Where("id=?",id).First(&bluetaber).Error

  if val != nil {
    utils.RespondWithError(w,http.StatusNotFound,"Bluetaber no encontrado")
    return
  } 

  json.NewEncoder(w).Encode(bluetaber)
}

