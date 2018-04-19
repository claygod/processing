package entities

// Processing
// Transaction
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"sort"
)

/*
Transactions types
*/

const (
	TransactionTypeTransfer int = iota
	TransactionTypeJoin
	TransactionTypeSeparation
	TransactionTypeExchange
)

/*
Transaction -
Транзакцию подписывает брокер (выходные блоки подписывает инициатор сделки).
Логика формирования выходов находится здесь.
*/
type Transaction struct {
	Initiator string
	Broker    string
	Type      int
	Inputs    []Block
	Outputs   []Block
	R         []byte
	S         []byte
	Hash      string
}

/*
NewTransaction - create new Transaction.
*/

func NewTransaction(initiator string, broker string, trType int) *Transaction {
	r := &Transaction{
		Initiator: initiator,
		Broker:    broker,
		Type:      trType,
		Inputs:    make([]Block, 0),
		Outputs:   make([]Block, 0),
	}
	return r
}

/*
AddInput - add Input.
*/
func (t *Transaction) AddInput(b Block) error {
	conditionCounter := 0
	if len(t.Outputs) > 0 {
		return fmt.Errorf("The outputs are already formed and inputs can not be added.")
	}
	for _, ib := range t.Inputs {
		if ib.Hash == b.Hash {
			return fmt.Errorf("Hash %s is already in the list of 'Inputs'.", ib.Hash)
		}
		if ib.Condition.ResourceId != 0 { // ToDo
			conditionCounter++
		}
	}
	if conditionCounter > 1 {
		return fmt.Errorf("It is impossible to combine several exchanges in one transaction.")
	}
	t.Inputs = append(t.Inputs, b)
	return nil
}

/*
Marshalling - preparation of data for hashing.
*/
func (t *Transaction) Marshalling() ([]byte, error) {
	t.R = []byte{}
	t.S = []byte{}
	t.Hash = "" //[]byte{}
	t.sortBlocks()

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(t)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

/*
MarshallingJson - preparation of data for hashing.
*/
func (t *Transaction) MarshallingJson() (string, error) {
	t.R = []byte{}
	t.S = []byte{}
	t.Hash = "" //[]byte{}
	t.sortBlocks()

	nb, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	//t.Hash = string(nb)
	return string(nb), nil
}

/*
SetHash - set hash.
*/
func (t *Transaction) SetHash(hash string) {
	t.Hash = hash
}

/*
sortBlocks - locks must be sorted for determinism.
*/
func (b *Transaction) sortBlocks() {
	sort.Slice(b.Inputs, func(i, j int) bool { return b.Inputs[i].Owner < b.Inputs[j].Owner })
	sort.Slice(b.Outputs, func(i, j int) bool { return b.Outputs[i].Owner < b.Outputs[j].Owner })
}

/*
checkSums - checking the sum of inputs and outputs.

func (b *Transaction) checkSums() {

}
*/
