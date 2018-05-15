package domain

// Processing
// Chain
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	// "fmt"
	"sync/atomic"
	"unsafe"
	// "github.com/claygod/processing/entities"
)

/*
Chain -
*/
type Chain struct {
	Counter         int64
	CurrentBlock    *Block
	OverlapBlock    *Block
	CurBlocks       *CurrentBlocks
	BlockRepository BlockRepository
}

/*
NewChain - create new Chain.
*/
func NewChain(bRepo BlockRepository) *Chain {
	ch := &Chain{
		CurrentBlock:    NewBlock(),
		OverlapBlock:    NewBlock(),
		CurBlocks:       NewCurrentBlocks(NewBlock(), NewBlock()),
		BlockRepository: bRepo,
	}
	return ch
}

func (c *Chain) AddTransaction(t *Transaction) {
	c.CurBlocks.CurrentBlock.WriteTransaction(t.Hash, t.Amount())
	c.CurBlocks.OverlapBlock.WriteTransaction(t.Hash, t.Amount())
}

func (c *Chain) GetCounter() int64 {
	return atomic.LoadInt64(&c.Counter)
}

func (c *Chain) SetCounter(newCount int64) {
	atomic.StoreInt64(&c.Counter, newCount)
}

func (c *Chain) Switch(t *Transaction) int64 {
	newBlock := NewBlock()
	nCb := NewCurrentBlocks(newBlock, c.CurrentBlock)
	addr := unsafe.Pointer(c.CurBlocks)
	atomic.StorePointer(&addr, unsafe.Pointer(nCb))
	c.BlockRepository.Write(atomic.LoadInt64(&c.Counter), c.OverlapBlock)
	return atomic.AddInt64(&c.Counter, 1)
}

/*
func (c *Chain) Verification(bh *BlockHash, num int64) error {
	b1, err := c.BlockRepository.Read(num - 1)
	if err != nil {
		return err
	}

	b2, err := c.BlockRepository.Read(num + 1)
	if err != nil {
		return err
	}
	h1 := b1.Transactions.Hashes()
	h2 := b2.Transactions.Hashes()
	for hash, _ := range bh.Transactions {
		if _, ok := h1[hash]; !ok {
			if _, ok := h2[hash]; !ok {
				fmt.Errorf("Transaction %s was not found.", hash)
			}
		}
	}
	return nil
}
*/
type CurrentBlocks struct {
	CurrentBlock *Block
	OverlapBlock *Block
}

func NewCurrentBlocks(cBlock *Block, oBlock *Block) *CurrentBlocks {
	c := &CurrentBlocks{
		CurrentBlock: cBlock,
		OverlapBlock: oBlock,
	}
	return c
}

func (c *CurrentBlocks) AddTransaction(t *Transaction) {
	c.CurrentBlock.WriteTransaction(t.Hash, t.Amount())
	c.OverlapBlock.WriteTransaction(t.Hash, t.Amount())
}
