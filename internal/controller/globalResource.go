package controller

import (
	"fmt"
	"k8s-mock/internal/categorizer"
	"k8s-mock/internal/dto"
	"k8s-mock/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func NewGlobalResourceController(repoResources *repository.Resource) *GlobalResource {
	return &GlobalResource{
		repoResources: repoResources,

		namespaceCategorizer: &categorizer.Namespace{},
	}
}

type GlobalResource struct {
	repoResources *repository.Resource

	namespaceCategorizer *categorizer.Namespace
}

func (ctrl *GlobalResource) Get(c *fiber.Ctx) error {
	var (
		rk dto.ResourceKey
	)
	err := makeInputBuilder(c).InURL(&rk).Error()
	if err != nil {
		return err
	}

	if ctrl.namespaceCategorizer.IsProject(&rk) {
		projects := ctrl.repoResources.Get(&rk)

		return c.Status(fiber.StatusOK).JSON(dto.Resource{
			"kind":       "ProjectList",
			"apiVersion": fmt.Sprintf("%s/%s", rk.APIGroup, rk.Version),
			"items":      projects,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.Resource{
		"kind":       "Status",
		"apiVersion": "v1",
		"metadata":   dto.Resource{},
		"status":     "Success",
	})
}

func (ctrl *GlobalResource) Create(c *fiber.Ctx) error {
	var (
		rk   dto.ResourceKey
		body dto.Resource
	)
	err := makeInputBuilder(c).InURL(&rk).InBody(&body).Error()
	if err != nil {
		return err
	}

	if ctrl.namespaceCategorizer.IsNamespace(&rk) ||
		ctrl.namespaceCategorizer.IsProject(&rk) {
		body["status"] = dto.Resource{"phase": "Active"}
		ctrl.repoResources.AppendNamespace(&body)
	} else {
		ctrl.repoResources.Append(&rk, &body)
	}
	return c.Status(fiber.StatusCreated).JSON(body)
}

func (ctrl *GlobalResource) GetNamespace(c *fiber.Ctx) error {
	var (
		rk dto.ResourceKey
	)
	err := makeInputBuilder(c).InURL(&rk).Error()
	if err != nil {
		return err
	}

	ns := ctrl.repoResources.GetNamespace(&rk)
	if ns == nil {
		return c.SendStatus(fiber.StatusNotFound)
	} else {
		return c.Status(fiber.StatusOK).JSON(ns)
	}
}
