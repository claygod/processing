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
Check - проверка заполненной транзакции.
*/
func (t *Transaction) Check(fee *Fee) error { // 1 fee = сотая часть процента
	switch t.Type {
	case TransactionTypeTransfer:
		t.checkTransfer(fee)
	case TransactionTypeJoin:

	case TransactionTypeSeparation:

	case TransactionTypeExchange:
	default:
		fmt.Errorf("Unsupported transaction type - %d.", t.Type)
	}
	return nil
}

/*
checkTransfer - проверка заполненной транзакции 'Transfer'.
*/
func (t *Transaction) checkTransfer(fee *Fee) error {
	// проверка на количество входов/выходов
	if len(t.Inputs) == 0 || len(t.Outputs) != 2 {
		fmt.Errorf("Few inputs (%d) and outputs (%d)", len(t.Inputs), len(t.Outputs))
	}
	// проверка на совпадение входных и выходных ресурсов
	if t.Inputs[0].State.ResourceId != t.Outputs[0].State.ResourceId {
		fmt.Errorf("Input resources (%d) do not coincide with output resources (%d).",
			t.Inputs[0].State.ResourceId, t.Outputs[0].State.ResourceId)
	}
	// проверка входов
	if err := t.checkTransferInputs(); err != nil {
		return err
	}
	// проверка выходов
	if err := t.checkTransferOutputs(fee); err != nil {
		return err
	}
	return nil
}

/*
checkTransferInputs - проверка входов заполненной транзакции 'Transfer'.
*/
func (t *Transaction) checkTransferInputs() error {
	resource := t.Inputs[0].State.ResourceId
	for _, b := range t.Inputs {
		if t.Initiator != b.Owner {
			return fmt.Errorf("Transfer: entries belong to more than one owner - %s, %s...", t.Initiator, b.Owner)
		}
		if resource != b.State.ResourceId {
			return fmt.Errorf("You can only transfer one resource at a time - %s, %s...", resource, b.State.ResourceId)
		}
		if b.Condition.Amount != 0 || b.Condition.ResourceId != 0 {
			return fmt.Errorf("In the transfer you can not use the inputs available for exchange.")
		}
	}
	return nil
}

/*
checkTransferOutputs - проверка выходов заполненной транзакции 'Transfer'.
*/
func (t *Transaction) checkTransferOutputs(fee *Fee) error {
	// проверка на совпадение входных и выходных ресурсов
	resource := t.Outputs[0].State.ResourceId
	for _, b := range t.Outputs {
		// совпадают ли брокеры транзакции и блоков
		if t.Broker != b.Broker {
			return fmt.Errorf("Brokers of transaction (%s) and block (%s) do not match.", t.Broker, b.Broker)
		}
		// выходные ресурсы должны быть одинаковыми
		if resource != b.State.ResourceId {
			return fmt.Errorf("You can only transfer one resource at a time - %s, %s...", resource, b.State.ResourceId)
		}
		// обменные условия должны быть обнулены
		if b.Condition.Amount != 0 || b.Condition.ResourceId != 0 {
			return fmt.Errorf("In the transfer you can not use the outputs available for exchange.")
		}
	}
	// поиск брокера и инициатора (отправителя)
	var brokerBlock *Block
	// var recipientBlock *Block
	if t.Outputs[0].Owner == t.Broker && t.Outputs[1].Owner == t.Initiator {
		brokerBlock = &t.Outputs[0]
		//recipientBlock = &t.Outputs[1]
	} else if t.Outputs[0].Owner == t.Broker && t.Outputs[0].Owner == t.Initiator {
		brokerBlock = &t.Outputs[1]
		//recipientBlock = &t.Outputs[0]
	} else {
		fmt.Errorf("Error in outputs.")
	}
	// проверка на общую сумму
	outAmount := t.Outputs[0].State.Amount + t.Outputs[1].State.Amount
	inAmount := t.inAmount()
	if outAmount != inAmount {
		fmt.Errorf("Do not match the amount of inputs (%d) and outputs (%d).", inAmount, outAmount)
	}
	// проверка на комиссионные брокеру
	feeAmount := fee.Count(outAmount)
	if feeAmount < 0 {
		fmt.Errorf("Error calculating transaction fees. (%d).", feeAmount)
	}
	// feeAmount := int(fee.Count(uint64(outAmount))) //t.toFee(outAmount, fee)
	if brokerBlock.State.Amount != feeAmount {
		fmt.Errorf("The expected brokerage fee is %d, and %d is indicated.", feeAmount, brokerBlock.State.Amount)
	}
	// проверка суммы получателю не делается, она косвенно в проверках на сумму и на комиссионные
	return nil
}

func (t *Transaction) inAmount() int {
	inAmount := 0
	for _, b := range t.Inputs {
		inAmount += b.State.Amount
	}
	return inAmount
}

func (t *Transaction) toFee(amount int, fee int) int {
	// amount * fee / 10000
	return 1
}

/*
Marshalling - preparation of data for hashing.
*/
func (t *Transaction) Marshalling() (string, error) {
	t.R = []byte{}
	t.S = []byte{}
	t.Hash = "" //[]byte{}
	t.sortBlocks()

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(t)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
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
