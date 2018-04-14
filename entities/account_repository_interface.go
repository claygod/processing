package entities

// Processing
// Account repository (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
AccountRepository - storage token interface.
This repository is not allowed to delete entities!
*/
type AccountRepository interface {
	Create([]byte) (string, error)
	Read(string) (Account, error) // by address
	List() []Account
}
