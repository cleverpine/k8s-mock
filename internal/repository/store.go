package repository

import (
	"k8s-mock/internal/dto"
	"sync"
)

func NewStoreRepository() *Store {
	return &Store{
		resources:  make(map[string][]dto.Resource),
		resourcesM: sync.Mutex{},
	}
}

type Store struct {
	resources  map[string][]dto.Resource
	resourcesM sync.Mutex
}

func (s *Store) Get(key string) []dto.Resource {
	s.resourcesM.Lock()
	defer s.resourcesM.Unlock()

	r, ok := s.resources[key]
	if ok {
		return r
	} else {
		tmp := []dto.Resource{}
		s.resources[key] = tmp
		return tmp
	}
}

func (s *Store) Set(key string, resources []dto.Resource) {
	s.resourcesM.Lock()
	defer s.resourcesM.Unlock()

	s.resources[key] = resources
}

func (s *Store) Append(key string, resource *dto.Resource) {
	s.resourcesM.Lock()
	defer s.resourcesM.Unlock()

	if resources, ok := s.resources[key]; ok {
		s.resources[key] = append(resources, *resource)
	} else {
		s.resources[key] = []dto.Resource{*resource}
	}
}

func (s *Store) Delete(key string, index int) {
	s.resourcesM.Lock()
	defer s.resourcesM.Unlock()

	if res, ok := s.resources[key]; ok {
		s.resources[key] = append(res[:index], res[(index+1):]...)
	}
}

func (s *Store) Find(key string, filter FilterFunc) (*dto.Resource, int) {
	s.resourcesM.Lock()
	defer s.resourcesM.Unlock()

	resources, ok := s.resources[key]

	if !ok {
		return nil, -1
	}

	for index, resource := range resources {
		if filter(&resource) {
			return &resource, index
		}
	}

	return nil, -1
}

func (s *Store) FindAll(key string, filter FilterFunc) []dto.Resource {
	s.resourcesM.Lock()
	defer s.resourcesM.Unlock()

	resources, ok := s.resources[key]

	if !ok {
		return nil
	}

	found := make([]dto.Resource, 0)
	for _, resource := range resources {
		if filter(&resource) {
			found = append(found, resource)
		}
	}

	return found
}

type FilterFunc func(r *dto.Resource) bool
