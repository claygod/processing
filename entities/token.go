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
func NewToken(pubKey []byte, enc Encoder) *Token {
	i := &Token{
		Address: enc.Encode(pubKey),
		PubKey:  pubKey,
	}
	return i
}
