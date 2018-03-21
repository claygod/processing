package processing

// Processing
// Recipient
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"math/big"
)

/*
Recipient -
*/
type Recipient struct {
	crypto *Crypto
}

/*
NewRecipient - create new Recipient.
*/

func NewRecipient(crypto *Crypto) *Recipient {
	r := &Recipient{
		crypto: crypto,
	}
	return r
}

func (r *Recipient) checkStamp(msg *Message, pubKey string) bool {
	rr, ok := new(big.Int).SetString(msg.Stamp.R10, 10)
	if !ok {
		return false
	}
	ss, ok := new(big.Int).SetString(msg.Stamp.S10, 10)
	if !ok {
		return false
	}

	return r.crypto.verify(
		msg.dataForVerification(),
		[]byte(pubKey),
		rr, ss)
}
