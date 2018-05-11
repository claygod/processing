package entities

// Processing
// Resource repository (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
ResourceRepository - storage token interface.
This repository is not allowed to delete entities!
*/
type ResourceRepository interface {
	Create(string) error
	Read(string) (*Resource, bool)
	Exists(string) bool
	// List() []*Resource
}
