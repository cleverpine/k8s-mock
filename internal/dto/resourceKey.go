package dto

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ResourceKey struct {
	NamespaceKey

	Namespace string
	Resource  string
}

func (rk *ResourceKey) Fill(c *fiber.Ctx) error {
	err := rk.NamespaceKey.Fill(c)
	if err != nil {
		return err
	}

	rk.Resource = c.Params("resource")

	return nil
}

func (rk *ResourceKey) UniqueKey() string {
	return fmt.Sprintf("%s/%s", rk.NamespaceKey.UniqueKey(), rk.Resource)
}
