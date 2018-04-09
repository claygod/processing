package entities

// Processing
// Token repository (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
TokenRepository - storage token interface.
*/
type TokenRepository interface {
	Create([]byte) error
	Read(string) (Token, error) // by address
	// Delete(string) error // токены не удаляются никогда!
}
