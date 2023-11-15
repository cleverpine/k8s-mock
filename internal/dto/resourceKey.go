package dto

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type ResourceKey struct {
	APIGroup  string `params:"apiGroup"`
	Version   string `params:"version"`
	Namespace string `params:"namespace"`
	Resource  string `params:"resource"`
}

func (rk *ResourceKey) Validate() error {
	rk.APIGroup = strings.ToLower(rk.APIGroup)
	rk.Version = strings.ToLower(rk.Version)
	rk.Namespace = strings.ToLower(rk.Namespace)
	rk.Resource = strings.ToLower(rk.Resource)

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
