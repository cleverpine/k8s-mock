package api

import (
	"k8s-mock/internal/controller"
	"k8s-mock/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func RESTV1(app *fiber.App) {
	var (
		repoResources = repository.NewResourceRepository()
	)

	var (
		debugCtrl         = controller.NewDebugController()
		apiDefinitionCtrl = controller.NewAPIDefinitionController()
		namespaceCtrl     = controller.NewNamespaceController(repoResources)
		metadataCtrl      = controller.NewMetadataController()

		globalResourceCtrl = controller.NewGlobalResourceController(repoResources)
		localResourceCtrl  = controller.NewLocalResourceController(repoResources)
	)

	app.Use(debugCtrl.Middleware)

	// app.Static("/", "./files")

	/**
		API DEFINITIONS
	**/
	app.Get("/version", metadataCtrl.Version)
	app.Get("/api", apiDefinitionCtrl.GetVersions)
	app.Get("/api/:version", apiDefinitionCtrl.GetAllAPIs)
	app.Get("/apis", apiDefinitionCtrl.GetAll)
	app.Get("/apis/:apiGroup/:version", apiDefinitionCtrl.Get)

	/**
		GLOBAL RESOURCE MANAGEMENT
	**/
	app.Get("/apis/:apiGroup/:version/:resourceType", globalResourceCtrl.Get)
	app.Post("/apis/:apiGroup/:version/:resourceType", globalResourceCtrl.Create)
	app.Get("/apis/:apiGroup/:version/:resourceType/~", globalResourceCtrl.GetUser)

	/**
		NAMESPACE MANAGEMENT
	**/
	app.Get("/api/:version/namespaces", namespaceCtrl.GetAll)
	app.Get("/api/:version/namespaces/:namespace", namespaceCtrl.Get)
	app.Get("/api/:version/projects/:namespace", namespaceCtrl.Get)
	app.Get("/apis/:apiGroup/:version/namespaces/:namespace", namespaceCtrl.Get)
	app.Get("/apis/:apiGroup/:version/projects/:namespace", namespaceCtrl.Get)

	app.Patch("/api/:version/namespaces/:namespace", namespaceCtrl.Update)
	app.Patch("/api/:version/projects/:namespace", namespaceCtrl.Update)
	app.Patch("/apis/:apiGroup/:version/namespaces/:namespace", namespaceCtrl.Update)
	app.Patch("/apis/:apiGroup/:version/projects/:namespace", namespaceCtrl.Update)

	app.Delete("/api/:version/namespaces/:namespace", namespaceCtrl.Delete)
	app.Delete("/api/:version/projects/:namespace", namespaceCtrl.Delete)
	app.Delete("/apis/:apiGroup/:version/namespaces/:namespace", namespaceCtrl.Delete)
	app.Delete("/apis/:apiGroup/:version/projects/:namespace", namespaceCtrl.Delete)

	/**
		LOCAL RESOURCE MANAGEMENT
	**/
	app.Get("/api/:version/namespaces/:namespace/:resourceType", localResourceCtrl.Get)
	app.Get("/apis/:apiGroup/:version/namespaces/:namespace/:resourceType", localResourceCtrl.Get)
	app.Get("/apis/:apiGroup/:version/projects/:namespace/:resourceType", localResourceCtrl.Get)

	app.Get("/apis/:apiGroup/:version/namespaces/:namespace/:resourceType/:resourceName", localResourceCtrl.GetSpecific)
	app.Get("/api/:version/namespaces/:namespace/:resourceType/:resourceName", localResourceCtrl.GetSpecific)

	app.Patch("/apis/:apiGroup/:version/namespaces/:namespace/:resourceType/:resourceName", localResourceCtrl.Update)
	app.Patch("/api/:version/namespaces/:namespace/:resourceType/:resourceName", localResourceCtrl.Update)

	app.Post("/api/:version/namespaces/:namespace/:resourceType", localResourceCtrl.Create)
	app.Post("/apis/:apiGroup/:version/namespaces/:namespace/:resourceType", localResourceCtrl.Create)
	app.Post("/apis/:apiGroup/:version/projects/:namespace/:resourceType", localResourceCtrl.Create)

	// app.Use(func(c *fiber.Ctx) error {
	// 	return c.SendStatus(fiber.StatusNotFound)
	// })

}
