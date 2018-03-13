package processing

// Processing
// Node test
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"testing"
)

// "n732h#$%%$DC" bjczMmgjJCUlJERD F9oJtbyc7t9mQGbQUWbw3yM39Hij34M5ZuhH3eaUvmbX
// "sdjh@#&gf@^*hg" c2RqaEAjJmdmQF4qaGc= DgkY5htw3KNH8Kgc7K4nwdJxaYETwXEoqoHfJG18veZY

func TestNodeNew(t *testing.T) { //t.Errorf("Encode(%s) = %s, want %s", pair.decoded, string(got), pair.encoded)
	n, err := NewNode("F9oJtbyc7t9mQGbQUWbw3yM39Hij34M5ZuhH3eaUvmbX", defaultAuthoritiesListPath)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("my URL: ", n.my.Url)
	}
	//fmt.Println("pvtkey: ", c.pvtKey.key())
	//fmt.Println("pubkey: ", c.pubKey)

}
