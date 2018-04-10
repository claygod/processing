package entities

// Processing
// Token repository (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
TokenRepository - storage token interface.
This repository is not allowed to delete entities!
*/
type TokenRepository interface {
	Create([]byte) (string, error)
	Read(string) (Token, error) // by address
	List() []Token
}
