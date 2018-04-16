package scripts

// Processing
// Encoder (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Encoder - the encoder interface that turns a set of bytes
into a human-readable string and vice versa.
*/
type Encoder interface {
	Encode([]byte) string
	Decode(string) []byte
	Address([]byte) string
}
