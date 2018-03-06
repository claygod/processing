package processing

// Processing
// Account
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	//"bytes"
	"crypto/sha256"
	// "strconv"
	"time"
	"unsafe"
)

const defaultBlockSize int = 128 // 120 + alignment

/*
Account - keeps a balance.
*/
type Account struct {
	Owner    int64
	Operator int64
	//counter  int64
	Balance int64
	//Reserve int64
	Chain []*Block
}

/*
NewAccount - create new account.
*/
func NewAccount(owner int64) *Account {
	a := &Account{
		Owner: owner,
		Chain: make([]*Block, 0),
	}
	return a
}

/*
Chain - keeps a blocks.
*/
type Chain struct {
	ch []*Block
}

func NewChain() *Chain {
	c := &Chain{ch: make([]*Block, defaultBlockSize)}
	c.AddBlock(NewBlock(0, 0, [32]byte{}))
	return c
}

func (c *Chain) AddBlock(b *Block) {
	c.ch = append(c.ch, b)
}

func (c *Chain) Check(shift int, h [32]byte) bool {
	if len(c.ch) < shift && c.ch[shift].Hash == h {
		return true
	}
	return false
}

type Block struct {
	Timestamp []byte // 40b
	Author    int64
	Addinion  int64
	PrevHash  [32]byte
	Hash      [32]byte
}

/*
NewBlock - create new block.
*/
func NewBlock(author int64, amount int64, prevHash [32]byte) *Block {
	b := &Block{
		Timestamp: []byte(time.Now().String()),
		Author:    author,
		Addinion:  amount,
		PrevHash:  prevHash,
	}
	b.Hash = hash(b)
	return b
}

/*
hash - calculate hash.
*/
func hash(b *Block) [32]byte {
	buf := make([]byte, 0)
	var b8 [8]byte
	b8 = *(*[8]byte)(unsafe.Pointer(&b.Author))
	buf = append(buf, b8[0:8]...)
	b8 = *(*[8]byte)(unsafe.Pointer(&b.Addinion))
	buf = append(buf, b8[0:8]...)
	//buf = append(buf, []byte(strconv.Itoa(b.Addinion))...)
	buf = append(buf, b.Timestamp...)
	buf = append(buf, b.PrevHash[0:32]...)
	return sha256.Sum256(buf)
}
