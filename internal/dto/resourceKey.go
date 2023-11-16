package dto

import (
	"fmt"
	"strings"
)

type ResourceKey struct {
	APIGroup     string `params:"apiGroup"`
	Version      string `params:"version"`
	Namespace    string `params:"namespace"`
	ResourceType string `params:"resourceType"`
	ResourceName string `params:"resourceName"`
}

func (rk *ResourceKey) Validate() error {
	rk.APIGroup = strings.ToLower(rk.APIGroup)
	rk.Version = strings.ToLower(rk.Version)
	rk.Namespace = strings.ToLower(rk.Namespace)
	rk.ResourceType = strings.ToLower(rk.ResourceType)

	// if rk.APIGroup == "" || rk.Version == "" {
	// 	return fiber.NewError(fiber.StatusBadRequest, "API Group and Version must be provided")
	// }

	return nil
}

func (rk *ResourceKey) Path() string {
	if rk.Namespace == "" {
		return fmt.Sprintf("%s/%s", rk.APIGroup, rk.Version)
	} else if rk.ResourceType == "" {
		return fmt.Sprintf("%s/%s/%s", rk.APIGroup, rk.Version, rk.Namespace)
	}

	return fmt.Sprintf("%s/%s/%s/%s", rk.APIGroup, rk.Version, rk.Namespace, rk.ResourceType)
}

func (rk *ResourceKey) IsK8sNamespace() bool {
	// TODO:
	return false
}

func (rk *ResourceKey) IsOSProject() bool {
	if rk.APIGroup == "project.openshift.io" {
		return rk.ResourceType == "projectrequest" || rk.ResourceType == "projectrequests" ||
			rk.ResourceType == "project" || rk.ResourceType == "projects"
	}

	return false
}
