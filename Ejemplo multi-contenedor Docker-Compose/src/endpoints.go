package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kerman-sanjuan/multi-contenedor/controllers"
)

func setupRoutes(app *fiber.App) {
    app.Get("/",controllers.Home)
	app.Post("/addBluetaber", controllers.CreateBluetaber)
	app.Get("/getBluetabers", controllers.GetBluetabers)
}