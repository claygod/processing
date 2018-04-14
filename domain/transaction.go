package domain

// Processing
// Transaction
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/claygod/processing/entities"
)

/*
Transaction -
Транзакцию подписывает брокер (выходные блоки подписывает инициатор сделки).
Логика формирования выходов находится здесь.
*/
type Transaction struct {
	Initiator string
	Broker    string
	Inputs    []entities.Block
	Outputs   []entities.Block
	R         []byte
	S         []byte
	Hash      string
}

/*
NewTransaction - create new Transaction.
*/

func NewTransaction(initiator string, broker string) *Transaction {
	r := &Transaction{
		Initiator: initiator,
		Broker:    broker,
		Inputs:    make([]entities.Block, 0),
		Outputs:   make([]entities.Block, 0),
	}
	return r
}

/*
AddInput - add Input.
*/
func (t *Transaction) AddInput(b entities.Block) error {
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
marshalling - preparation of data for hashing.
*/
func (t *Transaction) marshalling() (string, error) {
	t.R = []byte{}
	t.S = []byte{}
	t.Hash = "" //[]byte{}
	t.sortBlocks()

	nb, err := json.Marshal(t)
	if err != nil {
		return "", err
	}
	t.Hash = string(nb)
	return string(nb), nil
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
