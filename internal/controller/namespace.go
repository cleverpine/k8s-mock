package controller

import (
	"k8s-mock/internal/dto"
	"k8s-mock/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func NewNamespaceController(repo *repository.Namespace) *Namespace {
	return &Namespace{
		repo: repo,
	}
}

type Namespace struct {
	repo *repository.Namespace
}

func (ctrl *Namespace) GetAll(c *fiber.Ctx) error {
	nss := ctrl.repo.GetAll()

	return c.Status(fiber.StatusOK).JSON(dto.Resource{
		"apiVersion": "v1",
		"kind":       "NamespaceList",
		"items":      nss,
	})
}

func (ctrl *Namespace) Get(c *fiber.Ctx) error {
	var (
		rk dto.ResourceKey
	)
	err := makeInputBuilder(c).InURL(&rk).Error()
	if err != nil {
		return err
	}

	ns := ctrl.repo.Get(&rk)
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

	ns := ctrl.repo.Get(&rk)
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

	ns := ctrl.repo.Delete(&rk)
	if ns == nil {
		return c.SendStatus(fiber.StatusNotFound)
	} else {
		return c.Status(fiber.StatusOK).JSON(ns)
	}
}
