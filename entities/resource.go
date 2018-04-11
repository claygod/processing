package entities

// Processing
// Resource
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Resource - an important node in the network.
*/
type Resource struct {
	id   int
	name string
}

/*
NewResource - create new Resource.
*/
func NewResource(name string, id int) Resource {
	r := Resource{
		id:   id,
		name: name,
	}
	return r
}
