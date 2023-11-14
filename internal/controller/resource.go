package controller

import (
	"k8s-mock/internal/dto"

	"github.com/gofiber/fiber/v2"
)

func NewResourceController() *Resource {
	return &Resource{
		allocated: map[string][]dto.Resource{},
	}
}

type Resource struct {
	allocated map[string][]dto.Resource
}

func (ctrl *Resource) GetGlobal(c *fiber.Ctx) error {
	return ctrl.Get(c)
}

func (ctrl *Resource) Get(c *fiber.Ctx) error {
	var rk dto.ResourceKey
	if err := rk.Fill(c); err != nil {
		return err
	}

	tableResource := dto.Resource{
		"kind":       "Table",
		"apiVersion": "meta.k8s.io/v1",
		"columnDefinitions": []dto.Resource{
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

	resources, ok := ctrl.allocated[rk.UniqueKey()]
	if ok {
		tableResource["rows"] = resources
	} else {
		tableResource["rows"] = []dto.Resource{
			{
				"cells": []string{
					"a",
					"b",
					"c",
				},
				"object": dto.Resource{
					"kind":       "DeploymentConfig",
					"apiVersion": "apps.openshift.io/v1",
					"spec": map[string]interface{}{
						"replicas": 3,
					},
				},
			},
		}
	}

	return c.Status(fiber.StatusOK).JSON(tableResource)
}
