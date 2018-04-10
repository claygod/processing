package entities

// Processing
// Token
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Token - identifier.
*/
type Token struct {
	Address string
	PubKey  []byte
}

/*
NewToken - create new Token.
*/
func NewToken(address string, pubKey []byte) Token {
	t := Token{
		Address: address,
		PubKey:  pubKey,
	}
	return t
}
