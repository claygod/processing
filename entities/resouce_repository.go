package entities

// Processing
// Resources repository (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
ResourceRepository - storage token interface.
*/
type ResourceRepository interface {
	Create(string) (int, error)
	Read(string) (Resource, error) // by address
	GetNameById(string)
	Delete(string) error
}
