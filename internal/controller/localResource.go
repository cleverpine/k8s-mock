package controller

import (
	"k8s-mock/internal/dto"
	"k8s-mock/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func NewLocalResourceController(repoResources *repository.Resource) *LocalResource {
	return &LocalResource{
		repoResources: repoResources,
	}
}

type LocalResource struct {
	repoResources *repository.Resource
}

func (ctrl *LocalResource) Get(c *fiber.Ctx) error {
	var (
		rk dto.ResourceKey
	)
	err := makeInputBuilder(c).InURL(&rk).Error()
	if err != nil {
		return err
	}
	tableResource := dto.NewTableResource()

	resources := ctrl.repoResources.Get(&rk)
	if resources != nil {
		// TODO: use reflection to fill-in columnDefinitions
		tableResource.AddColumnDefinition("T", "string", "e", 0)
		tableResource.AddColumnDefinition("E", "string", "e", 0)
		tableResource.AddColumnDefinition("F", "string", "f", 0)

		tableResource.AddRow(
			[]string{"this", "is", "example"},
			&dto.GenericResource{
				"kind":       "DeploymentConfig",
				"apiVersion": "apps.openshift.io/v1",
				"spec": map[string]interface{}{
					"replicas": 3,
				},
			},
		)

	}

	return c.Status(fiber.StatusOK).JSON(tableResource)
}

func (ctrl *LocalResource) Create(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusInternalServerError)
}
