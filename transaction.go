package processing

// Processing
// Transaction
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"encoding/json"
	"sort"
)

/*
Transaction -
Транзакцию подписывает брокер (выходные блоки подписывает инициатор сделки).
*/
type Transaction struct {
	Initiator string
	Broker    string
	Inputs    []*Block
	Outputs   []*Block
	R         []byte
	S         []byte
	Hash      []byte
}

/*
NewTransaction - create new Transaction.
*/

func NewTransaction(initiator string, broker string) *Transaction {
	r := &Transaction{
		Initiator: initiator,
		Broker:    broker,
		Inputs:    make([]*Block, 0),
		Outputs:   make([]*Block, 0),
	}
	return r
}

/*
marshalling - preparation of data for hashing.
*/
func (t *Transaction) marshalling() ([]byte, error) {
	t.R = []byte{}
	t.S = []byte{}
	t.Hash = []byte{}
	t.sortBlocks()

	nb, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}
	return nb, nil
}

/*
sortBlocks - locks must be sorted for determinism.
*/
func (b *Transaction) sortBlocks() {
	sort.Slice(b.Inputs, func(i, j int) bool { return b.Inputs[i].OwnerAddress < b.Inputs[j].OwnerAddress })
	sort.Slice(b.Outputs, func(i, j int) bool { return b.Outputs[i].OwnerAddress < b.Outputs[j].OwnerAddress })
}

/*
checkSums - checking the sum of inputs and outputs.

func (b *Transaction) checkSums() {

}
*/
