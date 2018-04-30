package entities

// Processing
// Block
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	// "bytes"
	"fmt"
	"runtime"
	// "sort"
	"sync/atomic"
)

const i62 int64 = (1 << 63) - 1

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
	Transactions *TransactionsStore // []*Transaction
}

/*
NewBlock - create new Block.
*/
func NewBlock() *Block {
	b := &Block{
		//Number:       num,
		Transactions: NewTransactionsStore(), // make([]*Transaction, 0),
	}
	return b
}

func (b *Block) AddTransaction(t *Transaction) {
	if atomic.LoadInt64(&b.Saldo) > i62 {
		fmt.Errorf("The block is closed, you can not add it..")
	}
	b.Transactions.AddTransaction(t)
	/*
		b.lock()
		defer b.unlock()
		for _, s := range b.Transactions {
			if s.Hash == t.Hash {
				return
				// fmt.Errorf("Signature unit %s has already been added.", t.Hash)
			}
		}
		b.Transactions = append(b.Transactions, t)
	*/
	for _, o := range t.Body.Minus {
		atomic.AddInt64(&b.Saldo, o.Amount)
	}
	// sort.Slice(b.Transactions, func(i, j int) bool { return string(b.Transactions[i].Hash) < string(b.Transactions[j].Hash) })
	// return nil
}

func (b *Block) Close() { // (string, error)
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
