package entities

// Processing
// Account
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

/*
Account - identifier.
*/
type Account struct {
	Address string
	PubKey  []byte
}

/*
NewAccount - create new Account.
*/
func NewAccount(address string, pubKey []byte) Account {
	a := Account{
		Address: address,
		PubKey:  pubKey,
	}
	return a
}
