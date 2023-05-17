package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type bluetaber struct {
	ID        string `json:"id"`
	Nombre    string `json:"title"`
	Categoria string `json:"categoria"`
}

var bluetabers = []bluetaber{

	{ID: "1", Nombre: "Kerman", Categoria: "Technician"},
	{ID: "1", Nombre: "Tom Doe", Categoria: "Boss"},
	{ID: "1", Nombre: "John Doe", Categoria: "Technician"},
}

func getBluetabers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bluetabers)
}
	
func main(){
  router := gin.Default()
  router.GET("/bluetabers",getBluetabers)
  router.Run("0.0.0.0:8000")
}
