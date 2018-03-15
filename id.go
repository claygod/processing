package processing

// Processing
// Id
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Id - identifier.
*/
type Id struct {
	address string
	pubKey  []byte
}

/*
NewId - create new Id.
*/
func NewId(address string, pubKey []byte) *Id {
	i := &Id{
		address: address,
		pubKey:  pubKey,
	}
	return i
}
