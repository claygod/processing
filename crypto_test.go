package processing

// Processing
// Cripto test
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"testing"
)

func TestCriptoNew(t *testing.T) { //t.Errorf("Encode(%s) = %s, want %s", pair.decoded, string(got), pair.encoded)
	c, err := NewCrypto()
	if err != nil {
		t.Error("Error  - create new crypto.")
	}
	fmt.Println(c.pvtKey.key())
	fmt.Println(c.pubKey)
	fmt.Println(c.address)
}
