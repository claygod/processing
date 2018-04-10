package entities

// Processing
// Resources repository (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
ResourceRepository - storage token interface.
This repository is not allowed to delete entities!
*/
type ResourceRepository interface {
	Create(string) (int, error)
	Read(int) (Resource, error)
	// List () []Resource
}
