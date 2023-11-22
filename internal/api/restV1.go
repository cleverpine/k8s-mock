package api

import (
	"k8s-mock/internal/controller"
	"k8s-mock/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func RESTV1(app *fiber.App) {
	var (
		store = repository.NewStoreRepository()

		repoResource  = repository.NewResourceRepository(store)
		repoNamespace = repository.NewNamespaceRepository(store)
	)

	var (
		ctrlDebug         = controller.NewDebugController()
		ctrlApiDefinition = controller.NewAPIDefinitionController()
		ctrlNamespace     = controller.NewNamespaceController(repoNamespace)
		ctrlMetadata      = controller.NewMetadataController()

		ctrlGlobalResource = controller.NewGlobalResourceController(repoResource, repoNamespace)
		ctrlLocalResource  = controller.NewLocalResourceController(repoResource, repoNamespace)
	)

	app.Use(ctrlDebug.Middleware)

	// app.Static("/", "./files")

	/**
		API DEFINITIONS
	**/
	app.Get("/version", ctrlMetadata.Version)
	app.Get("/api", ctrlApiDefinition.GetAPI)
	app.Get("/api/:version", ctrlApiDefinition.GetAllAPIs)
	app.Get("/apis", ctrlApiDefinition.GetAll)
	app.Get("/apis/:apiGroup/:version", ctrlApiDefinition.Get)

	/**
		GLOBAL RESOURCE MANAGEMENT
	**/
	app.Get("/apis/:apiGroup/:version/:resourceType", ctrlGlobalResource.Get)
	app.Post("/apis/:apiGroup/:version/:resourceType", ctrlGlobalResource.Create)
	app.Get("/apis/:apiGroup/:version/:resourceType/~", ctrlGlobalResource.GetUser)

	/**
		NAMESPACE MANAGEMENT
	**/
	app.Get("/api/:version/namespaces", ctrlNamespace.GetAll)
	app.Get("/api/:version/namespaces/:namespace", ctrlNamespace.Get)
	app.Get("/api/:version/projects/:namespace", ctrlNamespace.Get)
	app.Get("/apis/:apiGroup/:version/namespaces/:namespace", ctrlNamespace.Get)
	app.Get("/apis/:apiGroup/:version/projects/:namespace", ctrlNamespace.Get)

	app.Patch("/api/:version/namespaces/:namespace", ctrlNamespace.Update)
	app.Patch("/api/:version/projects/:namespace", ctrlNamespace.Update)
	app.Patch("/apis/:apiGroup/:version/namespaces/:namespace", ctrlNamespace.Update)
	app.Patch("/apis/:apiGroup/:version/projects/:namespace", ctrlNamespace.Update)

	app.Delete("/api/:version/namespaces/:namespace", ctrlNamespace.Delete)
	app.Delete("/api/:version/projects/:namespace", ctrlNamespace.Delete)
	app.Delete("/apis/:apiGroup/:version/namespaces/:namespace", ctrlNamespace.Delete)
	app.Delete("/apis/:apiGroup/:version/projects/:namespace", ctrlNamespace.Delete)

	/**
		LOCAL RESOURCE MANAGEMENT
	**/
	app.Get("/api/:version/namespaces/:namespace/:resourceType", ctrlLocalResource.Get)
	app.Get("/apis/:apiGroup/:version/namespaces/:namespace/:resourceType", ctrlLocalResource.Get)
	app.Get("/apis/:apiGroup/:version/projects/:namespace/:resourceType", ctrlLocalResource.Get)

	app.Get("/apis/:apiGroup/:version/namespaces/:namespace/:resourceType/:resourceName", ctrlLocalResource.GetSpecific)
	app.Get("/api/:version/namespaces/:namespace/:resourceType/:resourceName", ctrlLocalResource.GetSpecific)

	app.Patch("/apis/:apiGroup/:version/namespaces/:namespace/:resourceType/:resourceName", ctrlLocalResource.Update)
	app.Patch("/api/:version/namespaces/:namespace/:resourceType/:resourceName", ctrlLocalResource.Update)

	app.Post("/api/:version/namespaces/:namespace/:resourceType", ctrlLocalResource.Create)
	app.Post("/apis/:apiGroup/:version/namespaces/:namespace/:resourceType", ctrlLocalResource.Create)
	app.Post("/apis/:apiGroup/:version/projects/:namespace/:resourceType", ctrlLocalResource.Create)

	// app.Use(func(c *fiber.Ctx) error {
	// 	return c.SendStatus(fiber.StatusNotFound)
	// })

}
