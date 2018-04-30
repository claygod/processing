package entities

// Processing
// Chain
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
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
	c.CurBlocks.CurrentBlock.AddTransaction(t)
	c.CurBlocks.OverlapBlock.AddTransaction(t)
}

func (c *Chain) Switch(t *Transaction) int64 {
	newBlock := NewBlock()
	oldBlock := c.OverlapBlock
	nCb := NewCurrentBlocks(newBlock, c.CurrentBlock)
	//addr := uintptr(unsafe.Pointer(c.CurBlocks))
	//atomic.StoreUintptr(&addr, uintptr(unsafe.Pointer(nCb)))
	addr := unsafe.Pointer(c.CurBlocks)
	atomic.StorePointer(&addr, unsafe.Pointer(nCb))
	oldBlock.Close()
	c.BlockRepository.Write(c.OverlapBlock)
	return 0
}

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
	c.CurrentBlock.AddTransaction(t)
	c.OverlapBlock.AddTransaction(t)
}
