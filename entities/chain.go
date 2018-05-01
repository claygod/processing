package entities

// Processing
// Chain
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	// "fmt"
	"sync/atomic"
	"unsafe"
)

/*
Chain -
*/
type Chain struct {
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

func (c *Chain) Switch(t *Transaction) int64 {
	newBlock := NewBlock()
	// oldBlock := c.OverlapBlock
	nCb := NewCurrentBlocks(newBlock, c.CurrentBlock)
	//addr := uintptr(unsafe.Pointer(c.CurBlocks))
	//atomic.StoreUintptr(&addr, uintptr(unsafe.Pointer(nCb)))
	addr := unsafe.Pointer(c.CurBlocks)
	atomic.StorePointer(&addr, unsafe.Pointer(nCb))
	//oldBlock.Close()
	c.BlockRepository.Write(c.OverlapBlock)
	return 0
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
