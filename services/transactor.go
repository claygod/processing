package services

// Processing
// Transactor
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"fmt"
	"sync/atomic"

	"github.com/claygod/processing/domain"
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
В параллельном режиме работает ТОЛЬКО метод Finish.
*/
type Transactor struct {
	status    int64
	accsMinus []*entities.Account
	accsPlus  []*entities.Account
	trn       *domain.Transaction
}

func NewTransactor() *Transactor {
	return &Transactor{
		status:    statusNov,
		accsMinus: make([]*entities.Account, 0),
		accsPlus:  make([]*entities.Account, 0),
	}
}

func (t *Transactor) Prepare(trn *domain.Transaction) error {
	if t.status != statusNov {
		return fmt.Errorf("This operation requires the status - %d, the actual status - %d.", statusNov, t.status)
	}
	t.trn = trn

	for i, op := range t.trn.Body.Minus {
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
		t.accsMinus[i].Unblock(t.trn.Body.Minus[i].Amount)
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
func (t *Transactor) Finish() (*domain.Transaction, error) {
	if !atomic.CompareAndSwapInt64(&t.status, statusPrepared, statusFinished) {
		return nil, fmt.Errorf("This operation requires the status - %d, the actual status - %d.", statusPrepared, t.status)
	}
	for i, op := range t.trn.Body.Minus {
		t.accsMinus[i].Credit(op.Amount)
	}
	for i, op := range t.trn.Body.Plus {
		t.accsPlus[i].Debit(op.Amount)
	}
	return t.trn, nil
}
