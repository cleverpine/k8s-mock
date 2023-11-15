package controller

import (
	"k8s-mock/internal/dto"
	"k8s-mock/internal/repository"

	"github.com/gofiber/fiber/v2"
)

func NewLocalResourceController(repoResources *repository.Resource) *LocalResource {
	return &LocalResource{
		repoResources: repoResources,
	}
}

type LocalResource struct {
	repoResources *repository.Resource
}

func (ctrl *LocalResource) Get(c *fiber.Ctx) error {
	var (
		rk dto.ResourceKey
	)
	err := makeInputBuilder(c).InURL(&rk).Error()
	if err != nil {
		return err
	}

	tableResource := dto.GenericResource{
		"kind":       "Table",
		"apiVersion": "meta.k8s.io/v1",
		"columnDefinitions": []dto.GenericResource{
			{
				"name":        "T",
				"type":        "string",
				"format":      "t",
				"description": "",
				"priority":    0,
			},
			{
				"name":        "E",
				"type":        "string",
				"format":      "e",
				"description": "",
				"priority":    0,
			},
			{
				"name":        "F",
				"type":        "string",
				"format":      "f",
				"description": "",
				"priority":    0,
			},
		},
	}

	resources := ctrl.repoResources.Get(&rk)
	if resources == nil {
		// TODO: use reflection to fill-in columnDefinitions
		tableResource["rows"] = []dto.GenericResource{
			{
				"cells": []string{
					"a",
					"b",
					"c",
				},
				"object": dto.GenericResource{
					"kind":       "DeploymentConfig",
					"apiVersion": "apps.openshift.io/v1",
					"spec": map[string]interface{}{
						"replicas": 3,
					},
				},
			},
		}
	} else {
		tableResource["rows"] = resources
	}

	return c.Status(fiber.StatusOK).JSON(tableResource)
}

func (ctrl *LocalResource) Create(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusInternalServerError)
}
