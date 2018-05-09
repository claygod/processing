package entities

// Processing
// Unit repository (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
UnitRepository - storage unit interface.
This repository is not allowed to delete entities!
*/
type UnitRepository interface {
	//Create([]byte) (string, error)
	Write(*Unit) error
	Read(string) (*Unit, error)
	//List() []Unit
}
