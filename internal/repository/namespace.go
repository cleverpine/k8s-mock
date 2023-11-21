package repository

import "k8s-mock/internal/dto"

func NewNamespaceRepository(s *Store) *Namespace {
	return &Namespace{
		s: s,
	}
}

type Namespace struct {
	s *Store
}

func (repo *Namespace) GetAll() []dto.Resource {
	return repo.s.Get("namespaces")
}

func (repo *Namespace) Get(key *dto.ResourceKey) *dto.Resource {
	r, _ := repo.get(key.Namespace)
	return r
}

func (repo *Namespace) Add(namespace *dto.Resource) {
	repo.s.Append("namespaces", namespace)
}

func (repo *Namespace) Delete(key *dto.ResourceKey) *dto.Resource {
	r, i := repo.get(key.Namespace)
	if i == -1 {
		return nil
	}

	repo.s.Delete("namespaces", i)
	return r
}

func (repo *Namespace) get(namespace string) (*dto.Resource, int) {
	return repo.s.Find(
		"namespaces",
		func(ns *dto.Resource) bool {
			return ns.GetString("metadata#name") == namespace
		},
	)
}
