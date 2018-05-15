package entities

// Processing
// Unit
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	// "fmt"
	"runtime"
	"sync/atomic"
	// "sync"
)

const maxKeyLen int = 100

const (
	unlocked int64 = iota
	blocked
)

/*
Unit - .

    A unit can be a customer, a company, etc.
    A unit can have many accounts (accounts are called a string variable)
    If a unit receives a certain amount for a nonexistent account, such an account will be created

*/
type Unit struct {
	hasp     int64
	pubKey   string
	accounts map[string]*Account
	// accsRepo AccountRepository
}

/*
NewUnit - create new Unit.
*/
func NewUnit() *Unit { // ar AccountRepository
	u := &Unit{
		accounts: make(map[string]*Account),
		//accsRepo: ar,
	}
	return u
}

func (u *Unit) Account(acc string) *Account {
	u.lock()
	defer u.unlock()
	a, ok := u.accounts[acc]
	if !ok {
		a = NewAccount() // u.getAccount(acc)
		u.accounts[acc] = a
	}
	return a
}

/*
func (u *Unit) getAccount(acc string) *Account {
	return u.accsRepo.Read(u.pubKey, acc)
}

func (u *Unit) Debit(acc string, amount int64) (int64, error) {
	if ln := len(acc); ln == 0 || ln > maxKeyLen {
		return -1, fmt.Errorf("Error in the length of the name of the requested account: %d. (Name: %s)", ln, acc)
	}
	return u.accsRepo.Read(u.pubKey, acc).Debit(amount)
}

func (u *Unit) Credit(acc string, amount int64) (int64, error) {
	if ln := len(acc); ln == 0 || ln > maxKeyLen {
		return -1, fmt.Errorf("Error in the length of the name of the requested account: %d. (Name: %s)", ln, acc)
	}
	return u.accsRepo.Read(u.pubKey, acc).Credit(amount)
}
*/
func (u *Unit) lock() {
	for {
		if atomic.CompareAndSwapInt64(&u.hasp, unlocked, blocked) {
			return
		}
		runtime.Gosched()
	}
}

func (u *Unit) unlock() {
	for {
		if atomic.CompareAndSwapInt64(&u.hasp, blocked, unlocked) {
			return
		}
		runtime.Gosched()
	}
}
