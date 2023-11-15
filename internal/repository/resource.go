package repository

import (
	"k8s-mock/internal/dto"
)

func NewResourceRepository() *Resource {
	return &Resource{
		resources: make(map[string][]dto.GenericResource),
		// namespaces: make(map[string]dto.GenericResource),
	}
}

type Resource struct {
	resources map[string][]dto.GenericResource
	// namespaces map[string]dto.GenericResource
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
	// TODO: add check for namespaces
	repo.resources["namespaces"] = append(repo.resources["namespaces"], *resource)
}

func (repo *Resource) GetNamespace(key *dto.ResourceKey) *dto.GenericResource {
	index := repo.getNamespaceIndex(key)
	if index == -1 {
		return nil
	}

	return &repo.resources["namespaces"][index]
}

func (repo *Resource) DeleteNamespace(key *dto.ResourceKey) *dto.GenericResource {
	index := repo.getNamespaceIndex(key)
	if index == -1 {
		return nil
	}

	nss := repo.resources["namespaces"]
	ns := nss[index]
	repo.resources["namespaces"] = append(nss[:index], nss[index+1:]...)

	return &ns
}

func (repo *Resource) getNamespaceIndex(key *dto.ResourceKey) int {
	for i, ns := range repo.resources["namespaces"] {
		if ns.Metadata.Name == key.Namespace {
			return i
		}
	}

	return -1
}
