package dto

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ResourceKey struct {
	APIGroup  string
	Version   string
	Namespace string
	Resource  string
}

func (rk *ResourceKey) Fill(c *fiber.Ctx) error {
	rk.APIGroup = c.Params("apiGroup")
	rk.Version = c.Params("version")
	rk.Namespace = c.Params("namespace")
	rk.Resource = c.Params("resource")

	if rk.APIGroup == "" || rk.Version == "" {
		return fiber.NewError(fiber.StatusBadRequest, "API Group and Version must be provided")
	}

	return nil
}

func (rk *ResourceKey) Path() string {
	if rk.Namespace == "" {
		return fmt.Sprintf("%s/%s", rk.APIGroup, rk.Version)
	} else if rk.Resource == "" {
		return fmt.Sprintf("%s/%s/%s", rk.APIGroup, rk.Version, rk.Namespace)
	}

	return fmt.Sprintf("%s/%s/%s/%s", rk.APIGroup, rk.Version, rk.Namespace, rk.Resource)
}
