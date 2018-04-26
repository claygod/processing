package entities

// Processing
// Unit
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	// "sync"
)

const maxKeyLen int = 100

/*
Unit - .

    A unit can be a customer, a company, etc.
    A unit can have many accounts (accounts are called a string variable)
    If a unit receives a certain amount for a nonexistent account, such an account will be created

*/
type Unit struct {
	// sync.RWMutex
	pubKey   []byte
	accsRepo AccountRepository
}

/*
NewUnit - create new Unit.
*/
func NewUnit(ar AccountRepository) *Unit {
	u := &Unit{
		accsRepo: ar,
	}
	return u
}

func (u *Unit) Debit(key string, amount int64) (int64, error) {
	if ln := len(key); ln == 0 || ln > maxKeyLen {
		return -1, fmt.Errorf("Error in the length of the name of the requested account: %d. (Name: %s)", ln, key)
	}
	return u.accsRepo.Read(key).Debit(amount)
}
