package repositories

// Processing
// Resource repository (implementation)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"sync"

	"github.com/claygod/processing/entities"
)

/*
ResourceRepository - storage resource (implementation).
This repository is not allowed to delete entities!
*/
type ResourceRepository struct {
	sync.RWMutex
	resources []entities.Resource
	// indexId   map[int]int
	indexName map[string]int
	// encoder   entities.Encoder
}

/*
NewResourceRepository - create new ResourceRepository.
*/
func NewResourceRepository() *ResourceRepository {
	r := &ResourceRepository{
		resources: make([]entities.Resource, 0),
		indexName: make(map[string]int),
	}
	return r
}

/*
Create - create new Resource.
Return resource-id.
*/
func (t *ResourceRepository) Create(name string) (int, error) {
	t.Lock()
	defer t.Unlock()
	if _, ok := t.indexName[name]; ok {
		return -1, fmt.Errorf("Resource %s already exists", name)
	}
	num := len(t.resources)
	nr := entities.NewResource(name, num)
	t.resources = append(t.resources, nr)
	t.indexName[name] = num
	return num, nil
}

/*
Read - get a resource at his id.
*/
func (t *ResourceRepository) Read(id int) (entities.Resource, error) {
	t.RLock()
	defer t.RUnlock()
	if len(t.resources) < id-1 {
		return entities.Resource{}, fmt.Errorf("Resource %d does not exist", id)
	}
	return t.resources[id], nil
}

/*
List - get a resources list.
*/
func (t *ResourceRepository) List() []entities.Resource {
	t.RLock()
	defer t.RUnlock()
	nrr := make([]entities.Resource, len(t.resources))
	copy(nrr, t.resources)
	return nrr
}
