package services

// Processing
// Transactor
// Copyright © 2018 Eduard Sesigin. All rights reserved. Contacts: <claygod@yandex.ru>

import (
	"github.com/claygod/processing/entities"
)

type Transactor struct {
	units *entities.UnitRepository
}

func NewTransactor(units *entities.UnitRepository) *Transactor {
	return &Transactor{
		units: units,
	}
}

func (m *Transactor) Prepare(t *entities.Transaction) error {
	return nil
}
func (m *Transactor) Rollback(t *entities.Transaction) error {
	return nil
}
func (m *Transactor) Execute(t *entities.Transaction) error {
	return nil
}
