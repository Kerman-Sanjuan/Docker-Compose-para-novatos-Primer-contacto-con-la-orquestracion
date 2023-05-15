package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kerman-sanjuan/multi-contenedor/models"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Bienvenido a la pagina de la API de Bluetab")
}

func CreateBluetaber(c *fiber.Ctx) error {
	fact := new(models.Bluetaber)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	models.DB.Create(&fact)

	return c.Status(200).JSON(fact)
}

func GetBluetabers(c *fiber.Ctx) error {
	bluetaber := []models.Bluetaber{}
	models.DB.Find(&bluetaber)

	return c.Status(200).JSON(bluetaber)
}
