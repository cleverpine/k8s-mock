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

func (ctrl *APIDefinition) GetVersions(c *fiber.Ctx) error {
	return sendFile(c, "files/api-versions.json")
}

func (ctrl *APIDefinition) GetAllAPIs(c *fiber.Ctx) error {
	var (
		rk dto.ResourceKey
	)
	err := makeInputBuilder(c).InURL(&rk).Error()
	if err != nil {
		return err
	}
	return sendFile(c, fmt.Sprintf("files/api-%s.json", rk.Version))
}

func (ctrl *APIDefinition) GetAll(c *fiber.Ctx) error {
	return sendFile(c, "files/apis.json")
}

func (ctrl *APIDefinition) Get(c *fiber.Ctx) error {
	var (
		rk dto.ResourceKey
	)
	err := makeInputBuilder(c).InURL(&rk).Error()
	if err != nil {
		return err
	}

	return sendFile(c, fmt.Sprintf("files/apis/%s/%s/serverresources.json", rk.APIGroup, rk.Version))
}
