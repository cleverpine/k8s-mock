package repository

import "k8s-mock/internal/dto"

func NewResourceRepository() *Resource {
	return &Resource{
		resources: make(map[string][]dto.Resource),
	}
}

type Resource struct {
	resources map[string][]dto.Resource
}

func (repo *Resource) Get(key *dto.ResourceKey) []dto.Resource {
	resources, ok := repo.resources[key.Path()]
	if ok {
		return resources
	} else {
		return resources
	}
}

func (repo *Resource) Replace(key *dto.ResourceKey, resources []dto.Resource) {
	repo.resources[key.Path()] = resources
}

func (repo *Resource) Append(key *dto.ResourceKey, resource *dto.Resource) {
	repo.resources[key.Path()] = append(repo.resources[key.Path()], *resource)
}

func (repo *Resource) AppendNamespace(resource *dto.Resource) {
	repo.resources["namespaces"] = append(repo.resources["namespaces"], *resource)
}

func (repo *Resource) GetNamespace(key *dto.ResourceKey) *dto.Resource {
	for _, ns := range repo.resources["namespaces"] {
		metadata, ok := ns["metadata"].(dto.Resource)
		if ok {
			if metadata["name"] == key.Namespace {
				return &ns
			}
		}
	}

	return nil
}
