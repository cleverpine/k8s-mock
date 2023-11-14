package controller

import (
	"fmt"
	"k8s-mock/internal/dto"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func NewResourceController() *Resource {
	return &Resource{
		allocated:  map[string][]dto.Resource{},
		namespaces: map[string]dto.Resource{},
	}
}

type Resource struct {
	allocated  map[string][]dto.Resource
	namespaces map[string]dto.Resource
}

func (ctrl *Resource) GetGlobal(c *fiber.Ctx) error {
	var rk dto.ResourceKey
	if err := rk.Fill(c); err != nil {
		return err
	}
	fmt.Println(1)
	fmt.Println(rk.APIGroup)
	fmt.Println(rk.Resource)
	if strings.ToLower(rk.APIGroup) == "project.openshift.io" &&
		strings.ToLower(rk.Resource) == "projects" {
		fmt.Println(2)
		ns := make([]dto.Resource, len(ctrl.namespaces))
		i := 0
		for _, r := range ctrl.namespaces {
			ns[i] = r
			i++
		}
		return c.Status(fiber.StatusOK).JSON(dto.Resource{
			"kind":       "ProjectList",
			"apiVersion": fmt.Sprintf("%s/%s", rk.APIGroup, rk.Version),
			"items":      ns,
		})
	}

	return c.Status(fiber.StatusOK).JSON(dto.Resource{
		"kind":       "Status",
		"apiVersion": "v1",
		"metadata":   dto.Resource{},
		"status":     "Success",
	})
}

func (ctrl *Resource) CreateGlobal(c *fiber.Ctx) error {
	var (
		rk   dto.ResourceKey
		body dto.Resource
	)
	if err := rk.Fill(c); err != nil {
		return err
	} else if err := c.BodyParser(&body); err != nil {
		return err
	}

	if strings.ToLower(rk.APIGroup) == "project.openshift.io" &&
		(strings.ToLower(rk.Resource) == "projectrequests" || strings.ToLower(rk.Resource) == "project") {
		rk.Resource = "Project"
		body["kind"] = "Project"
		body["status"] = dto.Resource{
			"phase": "Active",
		}
		md, ok := body["metadata"]
		if ok {
			ctrl.namespaces[md.(map[string]interface{})["name"].(string)] = body
		}
	}

	allocated, ok := ctrl.allocated[rk.Path()]
	if !ok {
		allocated = []dto.Resource{}
	}
	allocated = append(allocated, body)

	ctrl.allocated[rk.Path()] = allocated

	return c.Status(fiber.StatusCreated).JSON(body)
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

	resources, ok := ctrl.allocated[rk.Path()]
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

func (ctrl *Resource) Create(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusInternalServerError)
}

func (ctrl *Resource) GetNamespace(c *fiber.Ctx) error {
	var (
		rk dto.ResourceKey
	)
	if err := rk.Fill(c); err != nil {
		return err
	}

	ns, ok := ctrl.namespaces[rk.Namespace]
	if ok {
		return c.Status(fiber.StatusOK).JSON(ns)
	} else {
		return c.SendStatus(fiber.StatusNotFound)
	}
}
