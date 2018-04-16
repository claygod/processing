package scripts

// Processing
// Transaction interactor
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"github.com/claygod/processing/entities"
)

/*
TransactionInteractor - transactional use cases.
*/
type TransactionInteractor struct {
	authority entities.Authority
	eccoder   Encoder
}

/*
NewTransactionInteractor - create new TransactionInteractor.
*/
func NewTransactionInteractor() *TransactionInteractor {
	ti := &TransactionInteractor{}
	return ti
}

func (ti *TransactionInteractor) Transfer() { // t entities.Transaction

}
