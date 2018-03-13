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
func NewId(pubKey []byte) *Id {
	i := &Id{
		address: PubKeyToAddress(pubKey),
		pubKey:  pubKey,
	}
	return i
}
