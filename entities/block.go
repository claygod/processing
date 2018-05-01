package entities

// Processing
// Block
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	// "bytes"
	// "fmt"
	"runtime"
	// "sort"
	"sync/atomic"
)

const i62 int64 = (1 << 63) - 1 // ToDo: к удалению, т.к. закрывать блоки не будем

const (
	unlocked int64 = iota
	blocked
)

/*
Block -
*/
type Block struct {
	//Number       int64
	//Hash         string
	Saldo int64
	// Hasp         int64
	Transactions *TransactionsHashes
}

/*
NewBlock - create new Block.
*/
func NewBlock() *Block {
	b := &Block{
		Transactions: NewTransactionsHashes(),
	}
	return b
}

func (b *Block) WriteTransaction(hash string, amount int64) error {
	b.Transactions.AddTransaction(hash, amount)
	atomic.AddInt64(&b.Saldo, amount)
	return nil
}

func (b *Block) Hashes() map[string]int64 {
	return b.Transactions.HashesList()
}

func (b *Block) Close222() {
	//b.lock()
	//defer b.unlock()
	for {
		saldo := atomic.LoadInt64(&b.Saldo)
		if saldo > i62 {
			return
		}
		if atomic.CompareAndSwapInt64(&b.Saldo, saldo, saldo+i62) {
			return
		}
		runtime.Gosched()
	}

	/*
		sort.Slice(b.Transactions, func(i, j int) bool { return b.Transactions[i].Hash < b.Transactions[j].Hash })
		for {
			saldo := atomic.LoadInt64(&b.Saldo)
			if saldo > i62 {
				return
			}
			if atomic.CompareAndSwapInt64(&b.Saldo, saldo, saldo+i62) {
				return
			}
			runtime.Gosched()
		}
	*/
	/*
		var buf bytes.Buffer
		for _, t := range b.Transactions {
			_, err := buf.WriteString(t.Hash)
			if err != nil {
				return "", err
			}
		}
		b.Hash = buf.String()
	*/
	// return b.Hash, nil
}

/*
func (b *Block) lock() {
	for {
		if atomic.CompareAndSwapInt64(&b.Hasp, unlocked, blocked) {
			return
		}
		runtime.Gosched()
	}
}

func (b *Block) unlock() {
	for {
		if atomic.CompareAndSwapInt64(&b.Hasp, blocked, unlocked) {
			return
		}
		runtime.Gosched()
	}
}
*/

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

func (t *TransactionsHashes) AddTransaction(hash string, saldo int64) {
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
222 - К УДАЛЕНИЮ
*/

type TransactionsStore struct {
	Hasp         int64
	Transactions []*Transaction
}

func NewTransactionsStore() *TransactionsStore {
	b := &TransactionsStore{
		Hasp:         unlocked,
		Transactions: make([]*Transaction, 0),
	}
	return b
}

func (t *TransactionsStore) AddTransaction(tn *Transaction) {
	t.lock()
	defer t.unlock()
	t.Transactions = append(t.Transactions, tn)
}

/*
Hashes
Эта функция по идее будет использоваться редко,
но тем не менее реализована параллельно безопастно.
*/
func (t *TransactionsStore) Hashes() map[string]bool {
	out := make(map[string]bool)
	t.lock()
	defer t.unlock()
	for _, v := range t.Transactions {
		out[v.Hash] = true
	}
	return out
}

func (t *TransactionsStore) lock() {
	for {
		if atomic.CompareAndSwapInt64(&t.Hasp, unlocked, blocked) {
			return
		}
		runtime.Gosched()
	}
}

func (t *TransactionsStore) unlock() {
	for {
		if atomic.CompareAndSwapInt64(&t.Hasp, blocked, unlocked) {
			return
		}
		runtime.Gosched()
	}
}

type BlockHash struct {
	Saldo        int64
	Transactions map[string]bool
}
