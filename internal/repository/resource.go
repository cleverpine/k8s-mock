package repository

import (
	"k8s-mock/internal/dto"
)

func NewResourceRepository() *Resource {
	return &Resource{
		resources: make(map[string][]dto.GenericResource),
	}
}

type Resource struct {
	resources map[string][]dto.GenericResource
}

func (repo *Resource) Get(key *dto.ResourceKey) []dto.GenericResource {
	resources, ok := repo.resources[key.Path()]
	if ok {
		return resources
	} else {
		return resources
	}
}

func (repo *Resource) Replace(key *dto.ResourceKey, resources []dto.GenericResource) {
	repo.resources[key.Path()] = resources
}

func (repo *Resource) Append(key *dto.ResourceKey, resource *dto.GenericResource) {
	repo.resources[key.Path()] = append(repo.resources[key.Path()], *resource)
}

func (repo *Resource) GetNamespaces() []dto.GenericResource {
	ns, ok := repo.resources["namespaces"]
	if ok {
		return ns
	} else {
		return nil
	}
}

func (repo *Resource) AppendNamespace(resource *dto.GenericResource) {
	repo.resources["namespaces"] = append(repo.resources["namespaces"], *resource)
}

func (repo *Resource) GetNamespace(key *dto.ResourceKey) *dto.GenericResource {
	for _, ns := range repo.resources["namespaces"] {
		metadata, ok := ns["metadata"]
		if ok {
			metadataMap, ok := metadata.(map[string]any)
			if ok && metadataMap["name"] == key.Namespace {
				return &ns
			}
		}
	}

	return nil
}

func (repo *Resource) DeleteNamespace(key *dto.ResourceKey) *dto.GenericResource {
	nss := repo.resources["namespaces"]
	for indx, ns := range nss {
		metadata, ok := ns["metadata"]
		if ok {
			metadataMap, ok := metadata.(map[string]any)
			if ok && metadataMap["name"] == key.Namespace {
				repo.resources["namespaces"] = append(nss[:indx], nss[indx+1:]...)
				return &ns
			}
		}
	}

	return nil
}
