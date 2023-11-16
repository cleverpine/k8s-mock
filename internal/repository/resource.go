package repository

import (
	"k8s-mock/internal/dto"
)

func NewResourceRepository() *Resource {
	return &Resource{
		resources: make(map[string][]dto.Resource),
		// namespaces: make(map[string]dto.GenericResource),
	}
}

type Resource struct {
	resources map[string][]dto.Resource
	// namespaces map[string]dto.GenericResource
}

func (repo *Resource) Get(key *dto.ResourceKey) []dto.Resource {
	res := repo.resources[key.Path()]
	return res
}

func (repo *Resource) Replace(key *dto.ResourceKey, resources []dto.Resource) {
	repo.resources[key.Path()] = resources
}

func (repo *Resource) Append(key *dto.ResourceKey, resource *dto.Resource) {
	repo.resources[key.Path()] = append(repo.resources[key.Path()], *resource)
}

// func (repo *Resource) Delete(key *dto.ResourceKey) *dto.GenericResource {
// 	resources, ok := repo.resources[key.Path()]
// 	if ok {

// 	}
// 	delete(repo.resources, key.Path())
// }

func (repo *Resource) GetNamespaces() []dto.Resource {
	ns, ok := repo.resources["namespaces"]
	if ok {
		return ns
	} else {
		return nil
	}
}

func (repo *Resource) AppendNamespace(resource *dto.Resource) {
	// TODO: add check for namespace collision
	repo.resources["namespaces"] = append(repo.resources["namespaces"], *resource)
}

func (repo *Resource) GetNamespace(key *dto.ResourceKey) *dto.Resource {
	index := repo.getNamespaceIndex(key)
	if index == -1 {
		return nil
	}

	return &repo.resources["namespaces"][index]
}

func (repo *Resource) DeleteNamespace(key *dto.ResourceKey) *dto.Resource {
	index := repo.getNamespaceIndex(key)
	if index == -1 {
		return nil
	}

	nss := repo.resources["namespaces"]
	ns := nss[index]
	repo.resources["namespaces"] = append(nss[:index], nss[index+1:]...)

	return &ns
}

func (repo *Resource) FindResourceByFilter(key *dto.ResourceKey, filter FilterFunc) (*dto.Resource, int) {
	return repo.findResourceByFilter(key.Path(), filter)
}

func (repo *Resource) getNamespaceIndex(key *dto.ResourceKey) int {
	_, index := repo.findResourceByFilter(
		"namespaces",
		func(ns *dto.Resource) bool {
			return ns.GetString("metadata.name") == key.Namespace
		},
	)

	return index
}

func (repo *Resource) findResourceByFilter(key string, filter FilterFunc) (*dto.Resource, int) {
	for i, r := range repo.resources[key] {
		if filter(&r) {
			return &r, i
		}
	}
	return nil, -1
}

type FilterFunc func(r *dto.Resource) bool
