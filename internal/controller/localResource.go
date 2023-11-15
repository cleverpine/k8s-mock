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
				Name:     "API Group",
				Type:     "string",
				Priority: 0,
			},
			dto.ColumnDefinition{
				Name:     "Kind",
				Type:     "string",
				Priority: 0,
			},
			dto.ColumnDefinition{
				Name:     "Name",
				Type:     "string",
				Priority: 0,
			},
		)

		for _, res := range resources {
			tableResource.AddRows(
				dto.RowDefinition{
					Cells:  []string{res.APIVersion, res.Kind, res.Metadata.Name},
					Object: res,
				},
			)
		}
	}

	return c.Status(fiber.StatusOK).JSON(tableResource)
}

func (ctrl *LocalResource) Create(c *fiber.Ctx) error {
	var (
		rk   dto.ResourceKey
		body dto.GenericResource
	)
	err := makeInputBuilder(c).InURL(&rk).InBody(&body).Error()
	if err != nil {
		return err
	}

	ctrl.repoResources.Append(&rk, &body)
	return c.Status(fiber.StatusCreated).JSON(body)
}
