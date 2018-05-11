package services

// Processing
// Transactor
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"

	"github.com/claygod/processing/entities"
)

// Transactor status codes
const (
	statusNov int64 = iota
	statusPrepared
	statusFinished
)

/*
Transactor
Создаётся для каждой транзакции отдельно.
Только выполняет операции, записанные в транзакции.
Не предназначен для использования в параллельном режиме.
*/
type Transactor struct {
	status    int64
	accsMinus []*entities.Account
	accsPlus  []*entities.Account
	tBody     *entities.TransactionBody
}

func NewTransactor(accsMinus []*entities.Account, accsPlus []*entities.Account, tBody *entities.TransactionBody) *Transactor {
	return &Transactor{
		status:    statusNov,
		accsMinus: accsMinus,
		accsPlus:  accsPlus,
		tBody:     tBody,
	}
}

func (t *Transactor) Prepare() error {
	if t.status != statusNov {
		return fmt.Errorf("This operation requires the status - %d, the actual status - %d.", statusNov, t.status)
	}
	for i, op := range t.tBody.Minus {
		if _, err := t.accsMinus[i].Block(op.Amount); err != nil {
			t.rePrepare(i)
			return err
		}
	}
	t.status = statusPrepared
	return nil
}

func (t *Transactor) rePrepare(num int) {
	for i := 0; i < num; i++ {
		t.accsMinus[i].Unblock(t.tBody.Minus[i].Amount)
	}
}

func (t *Transactor) Rollback() error {
	if t.status != statusPrepared {
		return fmt.Errorf("This operation requires the status - %d, the actual status - %d.", statusPrepared, t.status)
	}
	t.rePrepare(len(t.accsMinus))
	t.status = statusNov
	return nil
}
func (t *Transactor) Finish() error {
	if t.status != statusPrepared {
		return fmt.Errorf("This operation requires the status - %d, the actual status - %d.", statusPrepared, t.status)
	}
	for i, op := range t.tBody.Minus {
		t.accsMinus[i].Credit(op.Amount)
	}
	for i, op := range t.tBody.Plus {
		t.accsPlus[i].Debit(op.Amount)
	}
	t.status = statusFinished
	return nil
}
