package controller

import (
	"k8s-mock/internal/dto"
	"k8s-mock/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func NewLocalResourceController(repoResource *repository.Resource, repoNamespace *repository.Namespace) *LocalResource {
	return &LocalResource{
		repoResource:  repoResource,
		repoNamespace: repoNamespace,
	}
}

type LocalResource struct {
	repoResource  *repository.Resource
	repoNamespace *repository.Namespace
}

func (ctrl *LocalResource) Get(c *fiber.Ctx) error {
	var (
		rk     dto.ResourceKey
		filter dto.ResourceFilter
	)
	err := makeInputBuilder(c).InURL(&rk).InQuery(&filter).Error()
	if err != nil {
		return err
	}

	var resources []dto.Resource

	if md := filter.GetMetadataFilter(); md != "" {
		r, _ := ctrl.repoResource.FindResourceByFilter(&rk, func(r *dto.Resource) bool {
			return r.GetString("metadata#name") == md &&
				r.GetString("metadata#namespace") == rk.Namespace
		})
		if r != nil {
			return c.Status(fiber.StatusOK).JSON(r)
		}
	} else {
		resources = ctrl.repoResource.FindResourcesByFilter(&rk, func(r *dto.Resource) bool {
			return r.GetString("metadata#namespace") == rk.Namespace
		})
	}

	// TODO: fix
	kind := "List"
	if rk.ResourceType == "secret" || rk.ResourceType == "secrets" {
		kind = "SecretList"
	} else if rk.ResourceType == "resourcequota" || rk.ResourceType == "resourcequotas" {
		kind = "ResourceQuotaList"
	}

	return c.Status(fiber.StatusOK).JSON(dto.Resource{
		"kind":       kind,
		"apiVersion": rk.Version,
		"items":      resources,
	})
}

func (ctrl *LocalResource) GetSpecific(c *fiber.Ctx) error {
	var (
		rk     dto.ResourceKey
		filter dto.ResourceFilter
	)
	err := makeInputBuilder(c).InURL(&rk).InQuery(&filter).Error()
	if err != nil {
		return err
	}
	r, _ := ctrl.repoResource.FindResourceByFilter(&rk, func(r *dto.Resource) bool {
		return r.GetString("metadata#name") == rk.ResourceName &&
			r.GetString("metadata#namespace") == rk.Namespace
	})

	if r == nil {
		return c.SendStatus(fiber.StatusNotFound)
	} else {
		return c.Status(fiber.StatusOK).JSON(r)
	}
}

func (ctrl *LocalResource) Create(c *fiber.Ctx) error {
	var (
		rk   dto.ResourceKey
		body dto.Resource
	)
	err := makeInputBuilder(c).InURL(&rk).InBody(&body).Error()
	if err != nil {
		return err
	}

	if r := ctrl.repoNamespace.Get(&rk); r == nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Error{Error: "Namespace doesn't exist"})
	}

	resourceName := body.GetString("metadata#name")
	r, _ := ctrl.repoResource.FindResourceByFilter(&rk, func(r *dto.Resource) bool {
		return r.GetString("metadata#name") == resourceName &&
			r.GetString("metadata#namespace") == rk.Namespace
	})

	if r != nil {
		return c.Status(fiber.StatusOK).JSON(r)
	}

	ctrl.repoResource.Append(&rk, &body)
	return c.Status(fiber.StatusCreated).JSON(body)
}

func (ctrl *LocalResource) Update(c *fiber.Ctx) error {
	var (
		rk   dto.ResourceKey
		body dto.Resource
	)
	err := makeInputBuilder(c).InURL(&rk).InBody(&body).Error()
	if err != nil {
		return err
	}

	if r := ctrl.repoNamespace.Get(&rk); r == nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.Error{Error: "Namespace doesn't exist"})
	}

	resource, _ := ctrl.repoResource.FindResourceByFilter(&rk, func(r *dto.Resource) bool {
		return r.GetString("metadata#name") == rk.ResourceName &&
			r.GetString("metadata#namespace") == rk.Namespace
	})

	if resource == nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	resource.Merge(&body)

	return c.Status(fiber.StatusOK).JSON(resource)
}
