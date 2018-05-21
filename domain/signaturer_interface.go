package domain

// Processing
// Signaturer (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Signaturer - signature service.
*/
type Signaturer interface {
	// Unit, Hash
	MakeSignature(string, string) *Signature
}
