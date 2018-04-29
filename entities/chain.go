package entities

// Processing
// Chain
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	// "fmt"
	"sync"
	"time"
)

/*
Chain -
*/
type Chain struct {
	sync.Mutex
	Counter      int64
	SwitchTime   int64
	CurrentBlock *Block
	OverlapBlock *Block
	Blocks       map[int64]*Block // ToDo: Repository
}

/*
NewChain - create new Chain.
*/
func NewChain(num int64) *Chain {
	if num == 0 {
		num = 1
	}
	ch := &Chain{
		Counter:      num,
		SwitchTime:   new(time.Time).Unix(),
		CurrentBlock: NewBlock(num),
		OverlapBlock: NewBlock(num - 1),
	}
	return ch
}

func (c *Chain) AddTransaction(t *Transaction) {
	c.Lock()
	defer c.Unlock()
	c.CurrentBlock.AddTransaction(t)
	c.OverlapBlock.AddTransaction(t)
}

func (c *Chain) Switch(t *Transaction) int64 {
	c.Lock()
	defer c.Unlock()
	c.SwitchTime = new(time.Time).Unix()
	nb := NewBlock(c.Counter + 1)
	c.OverlapBlock.Close()
	//if err != nil {
	//	return c.Counter, err
	//}
	c.Counter++
	c.CurrentBlock, c.OverlapBlock = nb, c.CurrentBlock
	c.addBlock(nb, c.Counter)
	return c.Counter
}

func (c *Chain) addBlock(b *Block, num int64) {
	c.Blocks[num] = b
}
