package entities

// Processing
// Provider
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"runtime"
	"sync/atomic"
)

/*
Provider - inputs and outputs resources from the network.
*/
type Provider struct {
	account  Account
	resource Resource
	limit    int64 // максимум, сколько можно вводить в систему
	counter  int64 // сколько введено в систему
}

/*
NewProvider - create new Provider.
*/
func NewProvider(id Account) *Provider {
	r := &Provider{
		account: id,
	}
	return r
}

func (p *Provider) GetLimit() int64 {
	return atomic.LoadInt64(&p.limit)
}

func (p *Provider) SetLimit(newLimit int64) bool {
	for {
		oldLimit := atomic.LoadInt64(&p.limit)
		if atomic.LoadInt64(&p.counter) > newLimit {
			return false
		} else if atomic.CompareAndSwapInt64(&p.limit, oldLimit, newLimit) {
			return true
		}
		runtime.Gosched()
	}
}

func (p *Provider) CurrentCount() int64 {
	return atomic.LoadInt64(&p.counter)
}

func (p *Provider) Addition(amount int64) int64 {
	for {
		limit := atomic.LoadInt64(&p.limit)
		oldCount := atomic.LoadInt64(&p.counter)
		newCount := oldCount + amount
		if newCount > limit && newCount < 0 {
			return -1
		} else if atomic.CompareAndSwapInt64(&p.counter, oldCount, newCount) {
			return newCount
		}
		runtime.Gosched()
	}
}
