package processing

// Processing
// Resource
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Resource - an important node in the network.
*/
type Resource struct {
	id   int64
	name string
}

/*
NewResource - create new Resource.
*/
func NewResource(name string) *Resource {
	r := &Resource{name: name}
	return r
}
