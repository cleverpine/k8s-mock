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
		tableResource.AddColumnDefinitions(
			dto.ColumnDefinition{
				Name:     "T",
				Type:     "string",
				Format:   "t",
				Priority: 0,
			},
			dto.ColumnDefinition{
				Name:     "E",
				Type:     "string",
				Format:   "e",
				Priority: 0,
			},
			dto.ColumnDefinition{
				Name:     "F",
				Type:     "string",
				Format:   "f",
				Priority: 0,
			},
		)

		tableResource.AddRows(
			dto.RowDefinition{
				Cells: []string{"this", "is", "example"},
				Object: &dto.GenericResource{
					Kind:       "DeploymentConfig",
					APIVersion: "apps.openshift.io/v1",
					Spec: dto.MapResource{
						"replicas": 3,
					},
				},
			},
		)

	}

	return c.Status(fiber.StatusOK).JSON(tableResource)
}

func (ctrl *LocalResource) Create(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusInternalServerError)
}
