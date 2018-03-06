package processing

// Processing
// Resource
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Resource - an important node in the network.
*/
type Resource struct {
	id int64
}

/*
NewResource - create new Resource.
*/
func NewResource() *Resource {
	a := &Resource{}
	return a
}
