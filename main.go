package main

import (
	"midtrans-example/controller"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	controller.Route(app)

	app.Listen(":3000")
}
