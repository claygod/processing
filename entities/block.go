package entities

// Processing
// Block
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"sort"
)

/*
Block -
*/
type Block struct {
	Number       int64
	Hash         string
	Transactions []*Transaction
}

/*
NewBlock - create new Block.
*/
func NewBlock(num int64) *Block {
	b := &Block{
		Number:       num,
		Transactions: make([]*Transaction, 0),
	}
	return b
}

func (b *Block) AddTransaction(t *Transaction) error {
	for _, s := range b.Transactions {
		if s.Hash == t.Hash {
			fmt.Errorf("Signature unit %s has already been added.", t.Hash)
		}
	}
	b.Transactions = append(b.Transactions, t)
	// sort.Slice(b.Transactions, func(i, j int) bool { return string(b.Transactions[i].Hash) < string(b.Transactions[j].Hash) })
	return nil
}

func (b *Block) Close() *Block {
	sort.Slice(b.Transactions, func(i, j int) bool { return b.Transactions[i].Hash < b.Transactions[j].Hash })
	return NewBlock(b.Number + 2)
}
