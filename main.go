package main

import (
	"k8s-mock/internal/api"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{})

	api.RESTV1(app)

	err := app.ListenTLS(":7777", "ssl/certificate.crt", "ssl/private-key.pem")
	panic(err.Error())
}
