package entities

// Processing
// Account
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"runtime"
	"sync/atomic"
)

/*
Account - keeps records of available and blocked funds.
Debit transaction is carried out in one step,
credit transaction in two steps (blocking and credit).
Operates in non-blocking mode.
*/
type Account struct {
	available int64
	blocked   int64
	//Address   string
	//PubKey    []byte
}

/*
NewAccount - create new Account.
*/
func NewAccount() Account {
	return Account{}
}

func (a *Account) Debit(amount int64) (int64, error) {
	if amount < 0 {
		return -1, fmt.Errorf("For a debit operation, the additive must be greater than zero: %d", amount)
	}
	res := atomic.AddInt64(&a.available, amount)
	if res < 0 {
		return res, fmt.Errorf("Debit operation result is below zero: %d", res)
	}
	return res, nil
}

func (a *Account) Block(amount int64) (int64, error) {
	if amount < 0 {
		return amount, fmt.Errorf("For the blocking operation, the digit must be greater than zero: %d", amount)
	}
	for {
		aviable := atomic.LoadInt64(&a.available)
		if aviable < amount {
			return -1, fmt.Errorf("Blocking error - there is %d, but blocked %d.", aviable, amount)
		}
		if atomic.CompareAndSwapInt64(&a.available, aviable, aviable-amount) {
			b := atomic.AddInt64(&a.blocked, amount)
			return b, nil
		}
		runtime.Gosched()
	}
}

func (a *Account) Credit(amount int64) (int64, error) {
	if amount < 0 {
		return amount, fmt.Errorf("For the credit operation, the digit must be greater than zero: %d", amount)
	}
	for {
		blocked := atomic.LoadInt64(&a.blocked)
		if blocked < amount {
			return -1, fmt.Errorf("Credit error - there is %d, but blocked %d.", amount, blocked)
		}
		if atomic.CompareAndSwapInt64(&a.blocked, blocked, blocked-amount) {
			b := atomic.AddInt64(&a.blocked, amount)
			return b, nil
		}
		runtime.Gosched()
	}
}
