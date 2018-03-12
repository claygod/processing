package processing

// Processing
// Node test
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"testing"
)

func TestNodeNew(t *testing.T) { //t.Errorf("Encode(%s) = %s, want %s", pair.decoded, string(got), pair.encoded)
	n, err := NewNode("8346834823yye23ye", defaultAuthoritiesListPath)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("my URL: ", n.my.Url)
	}
	//fmt.Println("pvtkey: ", c.pvtKey.key())
	//fmt.Println("pubkey: ", c.pubKey)

}
