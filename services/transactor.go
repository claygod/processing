package services

// Processing
// Transactor
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"github.com/claygod/processing/entities"
)

type Transactor struct {
	uRepo *entities.UnitRepository
	// tRepo *entities.TransactionRepository
}

func NewTransactor(units *entities.UnitRepository) *Transactor { // , transactions *entities.TransactionRepository
	return &Transactor{
		uRepo: units,
		// tRepo: transactions,
	}
}

func (m *Transactor) Prepare(t *entities.Transaction) error {
	return nil
}
func (m *Transactor) Rollback(key string) error {
	return nil
}
func (m *Transactor) Execute(key string) error {
	return nil
}
