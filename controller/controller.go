package controller

import (
	"midtrans-example/model"
	"midtrans-example/service"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Route(app *fiber.App) {
	app.Post("/create", Create)
	app.Get("/success", Success)
}

func Create(c *fiber.Ctx) error {
	var request model.MidtransRequest

	c.BodyParser(&request)

	res := service.Create(request)

	return c.Status(http.StatusOK).JSON(&model.WebResponse{
		Status:  "success",
		Message: "Transaction has been created",
		Data:    res,
	})
}

func Success(c *fiber.Ctx) error {
	return c.SendString("Payment is success")
}
