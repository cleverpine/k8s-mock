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

		return c.Status(fiber.StatusOK).JSON(dto.Resource{
			"apiVersion": fmt.Sprintf("%s/%s", rk.APIGroup, rk.Version),
			"kind":       "ProjectList",
			"items":      projects,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.Resource{
		"apiVersion": "v1",
		"kind":       "Status",
		"status": dto.Resource{
			"phase": "Success",
		},
	})
}

func (ctrl *GlobalResource) GetUser(c *fiber.Ctx) error {
	var (
		rk dto.ResourceKey
	)
	err := makeInputBuilder(c).InURL(&rk).Error()
	if err != nil {
		return err
	}

	// TODO: replace with actual user management
	return c.Status(fiber.StatusOK).JSON(dto.Resource{
		"kind":       "User",
		"apiVersion": fmt.Sprintf("%s/%s", rk.APIGroup, rk.Version),
		"metadata": dto.Resource{
			"name": "kubeadmin",
		},
		"identities": []string{
			"developer:kubeadmin",
		},
		"groups": []string{
			"system:authenticated",
			"system:authenticated:oauth",
		},
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

	if rk.IsK8sNamespace() || rk.IsOSProject() {
		body.Set("status", dto.Resource{"phase": "Active"})
		ctrl.repoResources.AppendNamespace(&body)
	} else {
		ctrl.repoResources.Append(&rk, &body)
	}
	return c.Status(fiber.StatusCreated).JSON(body)
}

// func (ctrl *GlobalResource) Delete(c *fiber.Ctx) error {
// 	var rk dto.ResourceKey
// 	err := makeInputBuilder(c).InURL(&rk).Error()
// 	if err != nil {
// 		return err
// 	}

// 	var resource *dto.GenericResource
// 	if rk.IsK8sNamespace() || rk.IsOSProject() {
// 		resource = ctrl.repoResources.DeleteNamespace(&rk)
// 	} else {
// 		resource = ctrl.repoResources.Delete(&rk)
// 	}
// 	return c.Status(fiber.StatusCreated).JSON(&resource)
// }
