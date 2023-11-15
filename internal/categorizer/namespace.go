package categorizer

import "k8s-mock/internal/dto"

type Namespace struct {
}

func (ns *Namespace) IsProject(rk *dto.ResourceKey) bool {
	if rk.APIGroup == "project.openshift.io" {
		return rk.Resource == "projectrequest" || rk.Resource == "projectrequests" || rk.Resource == "project" || rk.Resource == "projects"
	}

	return false
}

func (ns *Namespace) IsNamespace(rk *dto.ResourceKey) bool {
	// TODO:
	return false
}
