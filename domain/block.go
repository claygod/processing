package domain

// Processing
// Block
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	// "fmt"
	"runtime"
	"sync/atomic"
)

const (
	unlocked int64 = iota
	blocked
)

/*
Block -
*/
type Block struct {
	Saldo       int64
	HashesStore *TransactionsHashes
}

/*
NewBlock - create new Block.
*/
func NewBlock() *Block {
	b := &Block{
		HashesStore: NewTransactionsHashes(),
	}
	return b
}

func (b *Block) WriteTransaction(hash string, amount int64) error {
	b.HashesStore.AddHash(hash, amount)
	atomic.AddInt64(&b.Saldo, amount)
	return nil
}

func (b *Block) Hashes() map[string]int64 {
	return b.HashesStore.HashesList()
}

type TransactionsHashes struct {
	Hasp   int64
	Hashes map[string]int64
}

func NewTransactionsHashes() *TransactionsHashes {
	t := &TransactionsHashes{
		Hasp:   unlocked,
		Hashes: make(map[string]int64),
	}
	return t
}

func (t *TransactionsHashes) AddHash(hash string, saldo int64) {
	t.lock()
	defer t.unlock()
	t.Hashes[hash] = saldo
}

func (t *TransactionsHashes) HashesList() map[string]int64 {
	h := make(map[string]int64)
	t.lock()
	defer t.unlock()
	for k, v := range t.Hashes {
		h[k] = v
	}
	return h
}

func (t *TransactionsHashes) lock() {
	for {
		if atomic.CompareAndSwapInt64(&t.Hasp, unlocked, blocked) {
			return
		}
		runtime.Gosched()
	}
}

func (t *TransactionsHashes) unlock() {
	for {
		if atomic.CompareAndSwapInt64(&t.Hasp, blocked, unlocked) {
			return
		}
		runtime.Gosched()
	}
}

/*
BlockHash - служит для хранения ссылок на два актуальных блока
и возможности через пойнтер одновременно (атомарно) их сменить.

type BlockHash struct {
	Saldo        int64
	Transactions map[string]bool
}
*/
