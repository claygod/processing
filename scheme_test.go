package processing

// Processing
// Scheme test
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	// "fmt"
	"testing"
)

func TestSchemeListsLen(t *testing.T) { //t.Errorf("Encode(%s) = %s, want %s", string(got), pair.encoded)
	auths := make([]*Authority, 5, 5)

	s, err := NewScheme(auths, auths[0], 16)

	if err != nil {
		t.Error(err)
	}

	if len(s.list) != 10 {
		t.Error("Incorrect value len(s.list)")
	}
	//fmt.Println("pvtkey: ", c.pvtKey.key())
	//fmt.Println("pubkey: ", c.pubKey)

}

func TestSchemeGet(t *testing.T) { //t.Errorf("Encode(%s) = %s, want %s", string(got), pair.encoded)
	auths := make([]*Authority, 0, 51)
	for i := 0; i < 51; i++ {
		a := &Authority{}
		a.lastUpdate = int64(i)
		auths = append(auths, a)
	}

	s, err := NewScheme(auths, auths[0], 3)

	if err != nil {
		t.Error(err)
	}
	s.GetListToSend(0, 0)

}
