package entities

// Processing
// Resource
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Resource - an important node in the network.
*/
type Resource struct {
	id          string
	description string
}

/*
NewResource - create new Resource.
*/
func NewResource(id string, description string) *Resource {
	return &Resource{
		id:          id,
		description: description,
	}
}
