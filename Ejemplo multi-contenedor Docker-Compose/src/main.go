package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kerman-sanjuan/multi-contenedor/models"
) 

func main() {
    models.ConnectDB()
    app := fiber.New()

    setupRoutes(app)
    app.Listen(":3000")
}

