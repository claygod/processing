package services

// Processing
// Base58
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"math/big"
)

const alphabet string = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

type Base58 struct {
	alphabet  [58]byte
	decodeMap [256]int64
}

func NewBase58() *Base58 {

	enc := &Base58{}
	copy(enc.alphabet[:], []byte(alphabet)[:])
	for i := range enc.decodeMap {
		enc.decodeMap[i] = -1
	}
	for i, b := range enc.alphabet {
		enc.decodeMap[b] = int64(i)
	}
	return enc
}

func (b58 *Base58) Encode(key []byte) string {
	n := new(big.Int).SetBytes(key)
	addr, _ := b58.encodeNumber(n.String())
	return string(addr)
}

func (b58 *Base58) Decode(key string) []byte {
	addr, _ := b58.decodeBites([]byte(key))
	n, _ := new(big.Int).SetString(string(addr), 10)
	return n.Bytes()
}

func reverse(data []byte) {
	for i, j := 0, len(data)-1; i < j; i, j = i+1, j-1 {
		data[i], data[j] = data[j], data[i]
	}
}

// encodeNumber - encodes the number represented in the byte array base 10.
func (b58 *Base58) encodeNumber(src string) ([]byte, error) {
	if len(src) == 0 {
		return []byte{}, nil
	}
	n, ok := new(big.Int).SetString(src, 10)
	if !ok {
		return nil, fmt.Errorf("Expecting a number but got %q", src)
	}
	bytes := make([]byte, 0, len(src))
	for _, c := range src {
		if c == '0' {
			bytes = append(bytes, b58.alphabet[0])
		} else {
			break
		}
	}
	zerocnt := len(bytes)
	mod := new(big.Int)
	zero := big.NewInt(0)
	for {
		switch n.Cmp(zero) {
		case 1:
			n.DivMod(n, big.NewInt(58), mod)
			bytes = append(bytes, b58.alphabet[mod.Int64()])
		case 0:
			reverse(bytes[zerocnt:])
			return bytes, nil
		default:
			return nil, fmt.Errorf("Expecting a positive number in base58 encoding but got %q", n)
		}
	}
}

// decodeBites - decodes the base58 encoded bytes.
func (b58 *Base58) decodeBites(src []byte) ([]byte, error) {
	if len(src) == 0 {
		return []byte{}, nil
	}
	var zeros []byte
	for i, c := range src {
		if c == b58.alphabet[0] && i < len(src)-1 {
			zeros = append(zeros, '0')
		} else {
			break
		}
	}
	n := new(big.Int)
	var i int64
	for _, c := range src {
		if i = b58.decodeMap[c]; i < 0 {
			return nil, fmt.Errorf("Invalid character '%c' in decoding a base58 string \"%s\"", c, src)
		}
		n.Add(n.Mul(n, big.NewInt(58)), big.NewInt(i))
	}
	return n.Append(zeros, 10), nil
}
