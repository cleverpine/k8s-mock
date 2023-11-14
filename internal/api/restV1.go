package api

import (
	"k8s-mock/internal/controller"

	"github.com/gofiber/fiber/v2"
)

func RESTV1(app *fiber.App) {
	var (
		debugCtrl         = controller.NewDebugController()
		apiDefinitionCtrl = controller.NewAPIDefinitionController()
		resourceCtrl      = controller.NewResourceController()
	)

	app.Use(debugCtrl.Middleware)

	// app.Static("/", "./files")
	app.Get("/api", apiDefinitionCtrl.GetAllAPIs)
	app.Get("/apis", apiDefinitionCtrl.GetAll)
	app.Get("/apis/:apiGroup/:version", apiDefinitionCtrl.Get)

	app.Get("/apis/:apiGroup/:version/:resource", resourceCtrl.GetGlobal)
	app.Post("/apis/:apiGroup/:version/:resource", resourceCtrl.CreateGlobal)

	app.Get("/apis/:apiGroup/:version/namespaces/:namespace", resourceCtrl.GetNamespace)
	app.Get("/apis/:apiGroup/:version/projects/:namespace", resourceCtrl.GetNamespace)

	app.Get("/apis/:apiGroup/:version/namespaces/:namespace/:resource", resourceCtrl.Get)
	app.Get("/apis/:apiGroup/:version/projects/:namespace/:resource", resourceCtrl.Get)
	app.Post("/apis/:apiGroup/:version/namespaces/:namespace/:resource", resourceCtrl.Create)
	app.Post("/apis/:apiGroup/:version/projects/:namespace/:resource", resourceCtrl.Create)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

}
