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

	{ID: "1", Nombre: "Kerman Sanjuan", Categoria: "Technician"},
	{ID: "1", Nombre: "Tom", Categoria: "Boss"},
	{ID: "1", Nombre: "Kaiet Iglesias", Categoria: "Experienced Technician"},
}

func getBluetabers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bluetabers)
}
	
func main(){
  router := gin.Default()
  router.GET("/bluetabers",getBluetabers)
  router.Run("0.0.0.0:8000")
}
