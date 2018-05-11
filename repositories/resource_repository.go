package repositories

// Processing
// ResourceRepository
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"sync"

	"github.com/claygod/processing/entities"
)

/*
ResourceRepository - .
*/
type ResourceRepository struct {
	sync.RWMutex
	Data map[string]*entities.Resource
}

/*
NewResourceRepository - create new ResourceRepository.
*/
func NewResourceRepository() ResourceRepository {
	r := ResourceRepository{
		Data: make(map[string]*entities.Resource),
	}
	return r
}

func (r *ResourceRepository) Create(id string, description string) error {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.Data[id]; ok {
		return fmt.Errorf("Resource `%s` already exists.", id)
	}
	r.Data[id] = entities.NewResource(id, description)
	return nil
}

func (r *ResourceRepository) Read(id string) (*entities.Resource, bool) {
	r.RLock()
	defer r.RUnlock()
	res, ok := r.Data[id]
	return res, ok
}

func (r *ResourceRepository) Exists(id string) bool {
	r.RLock()
	defer r.RUnlock()
	_, ok := r.Data[id]
	return ok
}
