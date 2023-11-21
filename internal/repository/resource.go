package repository

import (
	"k8s-mock/internal/dto"
)

func NewResourceRepository(s *Store) *Resource {
	return &Resource{
		s: s,
	}
}

type Resource struct {
	s *Store
}

func (repo *Resource) Get(key *dto.ResourceKey) []dto.Resource {
	return repo.s.Get(key.Path())
}

func (repo *Resource) Append(key *dto.ResourceKey, newResource *dto.Resource) {
	var (
		k         = key.Path()
		resources = repo.s.Get(k)
	)

	newResourceName := newResource.GetString("metadata#name")
	newResourceNamespace := newResource.GetString("metadata#namespace")
	for _, resource := range resources {
		if resource.GetString("metadata#name") == newResourceName &&
			resource.GetString("metadata#namespace") == newResourceNamespace {
			return
		}
	}

	repo.s.Set(k, append(resources, *newResource))
}

// func (repo *Resource) Delete(key *dto.ResourceKey) *dto.Resource {
// 	resources, ok := repo.resources[key.Path()]
// 	if ok {

// 	}
// 	delete(repo.resources, key.Path())
// }

func (repo *Resource) FindResourcesByFilter(key *dto.ResourceKey, filter FilterFunc) []dto.Resource {
	return repo.s.FindAll(key.Path(), filter)
}

func (repo *Resource) FindResourceByFilter(key *dto.ResourceKey, filter FilterFunc) (*dto.Resource, int) {
	return repo.s.Find(key.Path(), filter)
}
