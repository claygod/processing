package services

// Processing
// Base58 test
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"testing"
)

type example struct {
	coder  *Base58
	tpairs []pair
}

type pair struct {
	decoded string
	encoded string
}

var examples = []example{
	{NewBase58(), []pair{
		{"", ""},
		{"0", "1"},
		{"32", "Z"},
		{"64", "27"},
		{"000", "111"},
		{"512", "9q"},
		{"1024", "Jf"},
		{"16777216", "2UzHM"},
		{"00068719476736", "1112ohWHHR"},
		{"18446744073709551616", "jpXCZedGfVR"},
		{"79228162514264337593543950336", "5qCHTcgbQwpvYZQ9d"},
	}},
}

func TestBase58Encode(t *testing.T) {
	for _, example := range examples {
		for _, pair := range example.tpairs {
			got, err := example.coder.encodeNumber(pair.decoded)
			if err != nil {
				t.Errorf("Error occurred while encoding %s.", pair.decoded)
			}
			if string(got) != pair.encoded {
				t.Errorf("Encode(%s) = %s, want %s", pair.decoded, string(got), pair.encoded)
			}
		}
	}
}

func TestBase58Decode(t *testing.T) {
	for _, example := range examples {
		for _, pair := range example.tpairs {
			got, err := example.coder.decodeBites([]byte(pair.encoded))
			if err != nil {
				t.Errorf("Error occurred while decoding %s.", pair.encoded)
			}
			if string(got) != pair.decoded {
				t.Errorf("Decode(%s) = %s, want %s", pair.encoded, string(got), pair.decoded)
			}
		}
	}
}

func BenchmarkBase58Encode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, example := range examples {
			for _, pair := range example.tpairs {
				_, _ = example.coder.encodeNumber(pair.decoded)
			}
		}
	}
}

func BenchmarkBase58Decode(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, example := range examples {
			for _, pair := range example.tpairs {
				_, _ = example.coder.decodeBites([]byte(pair.encoded))
			}
		}
	}
}
