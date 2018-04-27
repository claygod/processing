package entities

// Processing
// Transaction
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"bytes"
	"encoding/gob"
	"time"
)

/*
Transaction -
*/
type Transaction struct {
	Hash       string
	Body       *TransactionBody
	Signatures []*TransactionSignature
}

/*
NewTransaction - create new Transaction.
*/
func NewTransaction(initiator string, broker string) *Transaction {
	//t := &Transaction{}
	//t.newBody(initiator, broker)
	//return t
	return &Transaction{
		Body:       newTransactionBody(initiator, broker),
		Signatures: make([]*TransactionSignature, 0, 2),
	}
}

func (t *Transaction) Debit(unit string, account string, amount int64) *Transaction {
	t.Body.Plus = append(t.Body.Plus, TransactionOperation{
		Unit:    unit,
		Account: account,
		Amount:  amount,
	})
	return t
}

func (t *Transaction) Credit(unit string, account string, amount int64) *Transaction {
	t.Body.Plus = append(t.Body.Minus, TransactionOperation{
		Unit:    unit,
		Account: account,
		Amount:  amount,
	})
	return t
}

/*
func (t *Transaction) newBody(initiator string, broker string) {
	t.Body = &TransactionBody{
		Time:      new(time.Time).UnixNano(),
		Initiator: initiator,
		Broker:    broker,
		Plus:      make([]TransactionOperation, 0, 2),
		Minus:     make([]TransactionOperation, 0, 2),
	}
}
*/

type TransactionSignature struct {
	Unit string
	R    []byte
	S    []byte
}

type TransactionOperation struct {
	Unit    string
	Account string
	Amount  int64
}

type TransactionBody struct {
	Time int64
	// Initiator string
	Broker string
	Plus   []TransactionOperation
	Minus  []TransactionOperation
}

/*
Marshalling - preparation of data for hashing.
*/
func (t *TransactionBody) Marshalling() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(t)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func newTransactionBody(initiator string, broker string) *TransactionBody {
	return &TransactionBody{
		Time: new(time.Time).UnixNano(),
		// Initiator: initiator,
		Broker: broker,
		Plus:   make([]TransactionOperation, 0, 2),
		Minus:  make([]TransactionOperation, 0, 2),
	}
}
