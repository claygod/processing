package processing

// Processing
// Provider
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Provider - inputs and outputs resources from the network.
*/
type Provider struct {
	id    string
	limit int64
}

/*
NewProvider - create new Provider.
*/
func NewProvider() *Provider {
	r := &Provider{}
	return r
}
