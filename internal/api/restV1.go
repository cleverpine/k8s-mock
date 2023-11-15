package api

import (
	"k8s-mock/internal/controller"
	"k8s-mock/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func RESTV1(app *fiber.App) {
	var (
		resourceRepo = repository.NewResourceRepository()
	)

	var (
		debugCtrl         = controller.NewDebugController()
		apiDefinitionCtrl = controller.NewAPIDefinitionController()

		globalResourceCtrl = controller.NewGlobalResourceController(resourceRepo)
		localResourceCtrl  = controller.NewLocalResourceController(resourceRepo)
	)

	app.Use(debugCtrl.Middleware)

	// app.Static("/", "./files")
	app.Get("/api", apiDefinitionCtrl.GetAllAPIs)
	app.Get("/apis", apiDefinitionCtrl.GetAll)
	app.Get("/apis/:apiGroup/:version", apiDefinitionCtrl.Get)

	app.Get("/apis/:apiGroup/:version/:resource", globalResourceCtrl.Get)
	app.Post("/apis/:apiGroup/:version/:resource", globalResourceCtrl.Create)

	app.Get("/apis/:apiGroup/:version/namespaces/:namespace", globalResourceCtrl.GetNamespace)
	app.Get("/apis/:apiGroup/:version/projects/:namespace", globalResourceCtrl.GetNamespace)

	app.Get("/apis/:apiGroup/:version/namespaces/:namespace/:resource", localResourceCtrl.Get)
	app.Get("/apis/:apiGroup/:version/projects/:namespace/:resource", localResourceCtrl.Get)
	app.Post("/apis/:apiGroup/:version/namespaces/:namespace/:resource", localResourceCtrl.Create)
	app.Post("/apis/:apiGroup/:version/projects/:namespace/:resource", localResourceCtrl.Create)

	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

}
