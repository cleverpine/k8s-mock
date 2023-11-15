package controller

import (
	"fmt"
	"k8s-mock/internal/dto"
	"k8s-mock/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func NewGlobalResourceController(repoResources *repository.Resource) *GlobalResource {
	return &GlobalResource{
		repoResources: repoResources,
	}
}

type GlobalResource struct {
	repoResources *repository.Resource
}

func (ctrl *GlobalResource) Get(c *fiber.Ctx) error {
	var (
		rk dto.ResourceKey
	)
	err := makeInputBuilder(c).InURL(&rk).Error()
	if err != nil {
		return err
	}

	if rk.IsOSProject() {
		projects := ctrl.repoResources.GetNamespaces()

		return c.Status(fiber.StatusOK).JSON(dto.GenericResource{
			Kind:       "ProjectList",
			APIVersion: fmt.Sprintf("%s/%s", rk.APIGroup, rk.Version),
			Items:      projects,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.GenericResource{
		APIVersion: "v1",
		Kind:       "Status",
		Status: dto.ResourceStatus{
			Phase: "Success",
		},
	})
}

func (ctrl *GlobalResource) Create(c *fiber.Ctx) error {
	var (
		rk   dto.ResourceKey
		body dto.GenericResource
	)
	err := makeInputBuilder(c).InURL(&rk).InBody(&body).Error()
	if err != nil {
		return err
	}

	if rk.IsK8sNamespace() || rk.IsOSProject() {
		body.Status.Phase = "Active"
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

func (ctrl *GlobalResource) DeleteNamespace(c *fiber.Ctx) error {
	var (
		rk dto.ResourceKey
	)
	err := makeInputBuilder(c).InURL(&rk).Error()
	if err != nil {
		return err
	}

	ns := ctrl.repoResources.DeleteNamespace(&rk)
	if ns == nil {
		return c.SendStatus(fiber.StatusNotFound)
	} else {
		return c.Status(fiber.StatusOK).JSON(ns)
	}
}
