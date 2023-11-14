package dto

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type ResourceKey struct {
	Version string
	Name    string
}

func (rk *ResourceKey) Fill(c *fiber.Ctx) error {
	rk.Version = c.Params("version")
	rk.Name = c.Params("resource")

	return nil
}

func (rk *ResourceKey) UniqueKey() string {
	return fmt.Sprintf("%s/%s", rk.Version, rk.Name)
}
