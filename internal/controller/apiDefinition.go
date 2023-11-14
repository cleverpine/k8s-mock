package controller

import (
	"fmt"
	"k8s-mock/internal/dto"

	"github.com/gofiber/fiber/v2"
)

func NewAPIDefinitionController() *APIDefinition {
	return &APIDefinition{}
}

type APIDefinition struct {
}

func (ctrl *APIDefinition) GetAllAPIs(c *fiber.Ctx) error {
	return sendFile(c, "files/apis/v1/serverresources.json")
}

func (ctrl *APIDefinition) GetAll(c *fiber.Ctx) error {
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

func (ctrl *APIDefinition) Get(c *fiber.Ctx) error {
	var agk dto.APIGroupKey
	if err := agk.Fill(c); err != nil {
		return err
	}

	return sendFile(c, fmt.Sprintf("files/apis/%s/%s/serverresources.json", agk.APIGroup, agk.Version))
}
