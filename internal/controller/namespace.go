package controller

import (
	"k8s-mock/internal/dto"
	"k8s-mock/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func NewNamespaceController(repoResources *repository.Resource) *Namespace {
	return &Namespace{
		repoResources: repoResources,
	}
}

type Namespace struct {
	repoResources *repository.Resource
}

func (ctrl *Namespace) Get(c *fiber.Ctx) error {
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

func (ctrl *Namespace) Update(c *fiber.Ctx) error {
	var (
		rk   dto.ResourceKey
		body dto.Resource
	)
	err := makeInputBuilder(c).InURL(&rk).InBody(&body).Error()
	if err != nil {
		return err
	}

	ns := ctrl.repoResources.GetNamespace(&rk)
	if ns == nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	ns.Merge(&body)

	return c.Status(fiber.StatusOK).JSON(ns)
}

func (ctrl *Namespace) Delete(c *fiber.Ctx) error {
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
