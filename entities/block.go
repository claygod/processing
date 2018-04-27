package entities

// Processing
// Block
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
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

func (b *Block) Close() *Block {
	sort.Slice(b.Transactions, func(i, j int) bool { return b.Transactions[i].Hash < b.Transactions[j].Hash })
	return NewBlock(b.Number + 2)
}
