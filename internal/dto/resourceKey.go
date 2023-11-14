package dto

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ResourceKey struct {
	APIGroupKey

	Namespace string
	Resource  string
}

func (rk *ResourceKey) Fill(c *fiber.Ctx) error {
	err := rk.APIGroupKey.Fill(c)
	if err != nil {
		return err
	}

	rk.Namespace = c.Params("namespace", "") // namespace is optional for global resources
	rk.Resource = c.Params("resource")

	return nil
}

func (rk *ResourceKey) UniqueKey() string {
	return fmt.Sprintf("%s/%s/%s", rk.APIGroupKey.UniqueKey(), rk.Namespace, rk.Resource)
}
