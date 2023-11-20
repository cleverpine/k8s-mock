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

func (ctrl *LocalResource) GetTable(c *fiber.Ctx) error {
	var (
		rk dto.ResourceKey
	)
	err := makeInputBuilder(c).InURL(&rk).Error()
	if err != nil {
		return err
	}
	// tableResource := dto.NewTableResource()

	resources := ctrl.repoResources.FindResourcesByFilter(&rk, func(r *dto.Resource) bool {
		return r.GetString("metadata#namespace") == rk.Namespace
	})

	// if resources != nil {
	// 	// TODO: use reflection to fill-in columnDefinitions
	// 	tableResource.AddColumnDefinitions(
	// 		dto.ColumnDefinition{
	// 			Name:     "API Group",
	// 			Type:     "string",
	// 			Priority: 0,
	// 		},
	// 		dto.ColumnDefinition{
	// 			Name:     "Kind",
	// 			Type:     "string",
	// 			Priority: 0,
	// 		},
	// 		dto.ColumnDefinition{
	// 			Name:     "Name",
	// 			Type:     "string",
	// 			Priority: 0,
	// 		},
	// 	)

	// 	for _, res := range resources {
	// 		tableResource.AddRows(
	// 			dto.RowDefinition{
	// 				Cells:  []string{res.GetString("apiVersion"), res.GetString("kind"), res.GetString("metadata#name")},
	// 				Object: res,
	// 			},
	// 		)
	// 	}
	// }

	return c.Status(fiber.StatusOK).JSON(dto.Resource{
		"resources": resources,
	})
}

func (ctrl *LocalResource) GetSimple(c *fiber.Ctx) error {
	var (
		rk     dto.ResourceKey
		filter dto.ResourceFilter
	)
	err := makeInputBuilder(c).InURL(&rk).InQuery(&filter).Error()
	if err != nil {
		return err
	}

	var resources []dto.Resource

	if md := filter.GetMetadataFilter(); md != "" {
		r, _ := ctrl.repoResources.FindResourceByFilter(&rk, func(r *dto.Resource) bool {
			return r.GetString("metadata#name") == md &&
				r.GetString("metadata#namespace") == rk.Namespace
		})
		if r != nil {
			return c.Status(fiber.StatusOK).JSON(r)
		}
	} else {
		resources = ctrl.repoResources.Get(&rk)
	}

	// TODO: fix
	kind := "List"
	if rk.ResourceType == "secret" || rk.ResourceType == "secrets" {
		kind = "SecretList"
	} else if rk.ResourceType == "resourcequota" || rk.ResourceType == "resourcequotas" {
		kind = "ResourceQuotaList"
	}

	return c.Status(fiber.StatusOK).JSON(dto.Resource{
		"kind":       kind,
		"apiVersion": rk.Version,
		"items":      resources,
	})
}

func (ctrl *LocalResource) GetSpecific(c *fiber.Ctx) error {
	var (
		rk     dto.ResourceKey
		filter dto.ResourceFilter
	)
	err := makeInputBuilder(c).InURL(&rk).InQuery(&filter).Error()
	if err != nil {
		return err
	}
	r, _ := ctrl.repoResources.FindResourceByFilter(&rk, func(r *dto.Resource) bool {
		return r.GetString("metadata#name") == rk.ResourceName &&
			r.GetString("metadata#namespace") == rk.Namespace
	})

	if r == nil {
		return c.SendStatus(fiber.StatusNotFound)
	} else {
		return c.Status(fiber.StatusOK).JSON(r)
	}
}

func (ctrl *LocalResource) Create(c *fiber.Ctx) error {
	var (
		rk   dto.ResourceKey
		body dto.Resource
	)
	err := makeInputBuilder(c).InURL(&rk).InBody(&body).Error()
	if err != nil {
		return err
	}

	ctrl.repoResources.Append(&rk, &body)
	return c.Status(fiber.StatusCreated).JSON(body)
}

func (ctrl *LocalResource) Update(c *fiber.Ctx) error {
	var (
		rk   dto.ResourceKey
		body dto.Resource
	)
	err := makeInputBuilder(c).InURL(&rk).InBody(&body).Error()
	if err != nil {
		return err
	}

	resource, _ := ctrl.repoResources.FindResourceByFilter(&rk, func(r *dto.Resource) bool {
		return r.GetString("metadata#name") == rk.ResourceName &&
			r.GetString("metadata#namespace") == rk.Namespace
	})

	if resource == nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	resource.Merge(&body)

	return c.Status(fiber.StatusOK).JSON(body)
}
