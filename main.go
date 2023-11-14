package main

import (
	"k8s-mock/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var (
		debugCtrl    = controller.NewDebugController()
		apisCtrl     = controller.NewAPIsController()
		resourceCtrl = controller.NewResourceController()
	)

	app := fiber.New(fiber.Config{})

	app.Use(debugCtrl.Middleware)

	// app.Static("/", "./files")
	app.Get("/apis", apisCtrl.GetAll)
	app.Get("/apis/*", apisCtrl.Get)
	app.Get("/api/:vesion/:resource", resourceCtrl.Get)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	err := app.ListenTLS(":7777", "ssl/certificate.crt", "ssl/private-key.pem")
	panic(err.Error())
}
