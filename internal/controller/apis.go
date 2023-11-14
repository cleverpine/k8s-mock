package controller

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func NewAPIsController() *APIs {
	return &APIs{}
}

type APIs struct {
}

func (ctrl *APIs) GetAll(c *fiber.Ctx) error {
	// apis := dto.Resource{
	// 	"apiVersion": "v1",
	// 	"kind":       "APIResourceList",
	// 	"resources": []string{
	// 		"secrets",
	// 	},
	// }

	// return c.Status(fiber.StatusOK).JSON(apis)
	return sendFile(c, "files/apis.json")
}

func (ctrl *APIs) Get(c *fiber.Ctx) error {
	api := c.Params("*")

	return sendFile(c, fmt.Sprintf("files/apis/%s/serverresources.json", api))
}
