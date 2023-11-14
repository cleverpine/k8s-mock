package main

import (
	"k8s-mock/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func main() {
	var (
		debugCtrl         = controller.NewDebugController()
		apiDefinitionCtrl = controller.NewAPIDefinitionController()
		resourceCtrl      = controller.NewResourceController()
	)

	app := fiber.New(fiber.Config{})

	app.Use(debugCtrl.Middleware)

	// app.Static("/", "./files")
	app.Get("/api", apiDefinitionCtrl.GetAllAPIs)
	app.Get("/apis", apiDefinitionCtrl.GetAll)
	app.Get("/apis/:apiGroup/:version", apiDefinitionCtrl.Get)

	app.Get("/apis/:apiGroup/:vesion/:resource", resourceCtrl.GetGlobal)
	app.Get("/apis/:apiGroup/:vesion/namespaces/:namespace/:resource", resourceCtrl.Get)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	err := app.ListenTLS(":7777", "ssl/certificate.crt", "ssl/private-key.pem")
	panic(err.Error())
}
