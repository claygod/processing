package processing

// Processing
// Resource
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Resource - only the resources possible to transfer, exchange, sell.
*/
type Resource struct {
	limit int64
}

/*
NewResource - create new Resource.
*/
func NewResource() *Resource {
	r := &Resource{}
	return r
}
