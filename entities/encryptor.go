package entities

// Processing
// Encryptor (interface)
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"math/big"
)

/*
Encryptor - the encoder interface that turns a set of bytes
into a human-readable string and vice versa.
*/
type Encryptor interface {
	Sign([]byte) (*big.Int, *big.Int, error)
	Verify([]byte, []byte, *big.Int, *big.Int) bool
}
