package processing

// Processing
// Resource
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Resources - an important node in the network.
*/
type Resources struct {
	arr map[int]*Resource
}

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
func NewResource(name string) *Resource {
	r := &Resource{name: name}
	return r
}
