package controller

import (
	"k8s-mock/internal/dto"

	"github.com/gofiber/fiber/v2"
)

func NewResourceController() *Resource {
	return &Resource{
		allocated: map[string][]dto.Resource{},
	}
}

type Resource struct {
	allocated map[string][]dto.Resource
}

func (ctrl *Resource) Get(c *fiber.Ctx) error {
	var rk dto.ResourceKey
	if err := rk.Fill(c); err != nil {
		return err
	}

	resources, ok := ctrl.allocated[rk.UniqueKey()]
	if ok {
		return c.Status(fiber.StatusOK).JSON(resources)
	}

	return c.Status(fiber.StatusOK).JSON([]dto.Resource{})
}
