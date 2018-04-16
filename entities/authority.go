package entities

// Processing
// Authority
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"sync/atomic"
)

/*
Authority - an important node in the network.
*/
type Authority struct {
	Account Account
	Link    string `json:"url"`
	Status  int64
}

/*
NewAuthority - create new Authority.
*/
func NewAuthority(account Account, link string) Authority {
	a := Authority{
		Account: account,
		Link:    link,
	}
	return a
}

func (a *Authority) StatusAdd(amount int64) int64 { // ToDo: the "method" mov to Node
	return atomic.AddInt64(&a.Status, amount)
}
